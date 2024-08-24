package repository

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/yogarn/arten/entity"
)

type ISessionRepository interface {
	CreateSession(session *entity.Session) error
	CheckSession(refreshToken string) error
	DeleteSession(userId uuid.UUID, refreshToken string) error
}

type SessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) ISessionRepository {
	return &SessionRepository{db}
}

func (sessionRepository *SessionRepository) CreateSession(session *entity.Session) error {
	stmt := `
		INSERT INTO sessions (id, user_id, refresh_token, device_info, ip_address, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	tx, err := sessionRepository.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(stmt, session.Id, session.UserId, session.RefreshToken, session.DeviceInfo, session.IpAddress, session.CreatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

func (sessionRepository *SessionRepository) CheckSession(refreshToken string) error {
	stmt := `
		SELECT * FROM sessions WHERE refresh_token = ?
	`

	row := sessionRepository.db.QueryRow(stmt, refreshToken)
	session := &entity.Session{}
	err := row.Scan(&session.Id, &session.UserId, &session.RefreshToken, &session.DeviceInfo, &session.IpAddress, &session.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (sessionRepository *SessionRepository) DeleteSession(userId uuid.UUID, refreshToken string) error {
	stmt := `
		DELETE FROM sessions WHERE user_id = ? AND refresh_token = ?
	`

	tx, err := sessionRepository.db.Begin()
	if err != nil {
		return err
	}

	result, err := tx.Exec(stmt, userId, refreshToken)
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
		return errors.New("no session deleted")
	}

	err = tx.Commit()
	return err
}
