package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/yogarn/arten/entity"
	"github.com/yogarn/arten/model"
)

type IUserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	LoginUser(user *model.UserLogin) (*entity.User, error)
	GetUserById(id uuid.UUID) (*entity.User, error)
	UpdateUser(id uuid.UUID, userReq *model.UpdateUser) (*model.UpdateUser, error)
	SendOtp(username string) error
	CheckOtp(otpRequest *model.OtpRequest) error
	ActivateUser(username string) error
}

type UserRepository struct {
	db    *sql.DB
	redis *redis.Client
}

func NewUserRepository(db *sql.DB, redis *redis.Client) IUserRepository {
	return &UserRepository{
		db:    db,
		redis: redis,
	}
}

func (userRepository *UserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	stmt := `INSERT INTO users (id, name, username, password, email, is_verified, created_at) 
	VALUES (?, ?, ?, ?, ?, FALSE, UTC_TIMESTAMP())`

	tx, err := userRepository.db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(stmt, user.ID, user.Name, user.Username, user.Password, user.Email)
	if err != nil {
		tx.Rollback()
		return user, err
	}

	err = tx.Commit()
	return user, err
}

func (userRepository *UserRepository) SendOtp(username string) error {
	stmt := `SELECT * FROM users WHERE username = ?`
	tx, err := userRepository.db.Begin()
	if err != nil {
		return err
	}

	user := &entity.User{}

	row := tx.QueryRow(stmt, username)
	err = row.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Email, &user.IsVerified, &user.CreatedAt)

	if err != nil {
		return err
	}

	otp := fmt.Sprintf("%04d", rand.Intn(9000)+1000)
	redisOtp := userRepository.redis.Set(context.Background(), user.Email, otp, 5*time.Minute)

	if redisOtp.Err() != nil {
		return errors.New("failed to send otp")
	}

	return nil
}

func (userRepository *UserRepository) CheckOtp(otpRequest *model.OtpRequest) error {
	stmt := `SELECT * FROM users WHERE username = ?`
	tx, err := userRepository.db.Begin()
	if err != nil {
		return err
	}

	user := &entity.User{}

	row := tx.QueryRow(stmt, otpRequest.Username)
	err = row.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Email, &user.IsVerified, &user.CreatedAt)

	if err != nil {
		return err
	}

	redisOtp := userRepository.redis.Get(context.Background(), user.Email)
	if redisOtp.Err() != nil {
		return errors.New("failed to use redis")
	}

	result, err := redisOtp.Result()
	if err != nil {
		return errors.New("failed to get otp")
	}

	if otpRequest.Otp != result {
		return errors.New("invalid otp")
	}

	return nil
}

func (userRepository *UserRepository) ActivateUser(username string) error {
	stmt := "UPDATE users SET is_verified = TRUE WHERE username = ?"

	tx, err := userRepository.db.Begin()
	if err != nil {
		return err
	}

	result, err := tx.Exec(stmt, username)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if rowsAffected <= 0 {
		tx.Rollback()
		return errors.New("no row updated")
	}

	err = tx.Commit()
	return err
}

func (userRepository *UserRepository) LoginUser(userReq *model.UserLogin) (*entity.User, error) {
	stmt := `SELECT * FROM users WHERE username = ?`
	tx, err := userRepository.db.Begin()
	if err != nil {
		return nil, err
	}

	user := &entity.User{}

	row := tx.QueryRow(stmt, userReq.Username)
	err = row.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Email, &user.IsVerified, &user.CreatedAt)

	if !user.IsVerified {
		return nil, errors.New("user not verified")
	}

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	return user, err
}

func (userRepository *UserRepository) GetUserById(id uuid.UUID) (*entity.User, error) {
	stmt := `SELECT * FROM users WHERE id = ?`
	tx, err := userRepository.db.Begin()
	if err != nil {
		return nil, err
	}

	user := &entity.User{}

	row := tx.QueryRow(stmt, id)
	err = row.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Email, &user.IsVerified, &user.CreatedAt)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	return user, err
}

func (userRepository *UserRepository) UpdateUser(id uuid.UUID, userReq *model.UpdateUser) (*model.UpdateUser, error) {
	var column []string
	var values []interface{}

	if userReq.Name != "" {
		column = append(column, "name = ?")
		values = append(values, userReq.Name)
	}
	if userReq.Username != "" {
		column = append(column, "username = ?")
		values = append(values, userReq.Username)
	}
	if userReq.Password != "" {
		column = append(column, "password = ?")
		values = append(values, userReq.Password)
	}
	if userReq.Email != "" {
		column = append(column, "email = ?")
		values = append(values, userReq.Email)
	}

	values = append(values, id)

	stmt := fmt.Sprintf("UPDATE users SET %s WHERE id = ?", strings.Join(column, ", "))

	tx, err := userRepository.db.Begin()
	if err != nil {
		return nil, err
	}

	result, err := tx.Exec(stmt, values...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if rowsAffected <= 0 {
		tx.Rollback()
		return nil, errors.New("no row updated")
	}

	err = tx.Commit()
	return userReq, err
}
