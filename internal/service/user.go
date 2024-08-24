package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yogarn/arten/entity"
	"github.com/yogarn/arten/internal/repository"
	"github.com/yogarn/arten/model"
	"github.com/yogarn/arten/pkg/bcrypt"
	"github.com/yogarn/arten/pkg/jwt"
	"github.com/yogarn/arten/pkg/smtp"
)

type IUserService interface {
	Register(userReq *model.UserRegister) (*entity.User, error)
	Login(userReq *model.UserLogin) (*model.UserLoginResponse, error)
	RefreshToken(token string) (string, error)
	GetUserById(id uuid.UUID) (*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
	UpdateUser(ctx *gin.Context, user *model.UpdateUser) (*model.UpdateUser, error)
	SendOtp(username string) error
	ActivateUser(otpRequest *model.OtpRequest) error
	GetIpAddress(ctx *gin.Context) string
	GetDeviceInfo(ctx *gin.Context) string
}

type UserService struct {
	UserRepository repository.IUserRepository
	Bcrypt         bcrypt.Interface
	JWT            jwt.Interface
	SMTP           smtp.Interface
}

func NewUserService(userRepository repository.IUserRepository, bcrypt bcrypt.Interface, jwt jwt.Interface, smtp smtp.Interface) IUserService {
	return &UserService{
		UserRepository: userRepository,
		Bcrypt:         bcrypt,
		JWT:            jwt,
		SMTP:           smtp,
	}
}

func (userService *UserService) Register(userReq *model.UserRegister) (*entity.User, error) {
	hashPassword, err := userService.Bcrypt.GenerateFromPassword(userReq.Password)
	if err != nil {
		return nil, err
	}
	userEntity := &entity.User{
		ID:        uuid.New(),
		Username:  userReq.Username,
		Password:  hashPassword,
		Name:      userReq.Name,
		Email:     userReq.Email,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	user, err := userService.UserRepository.CreateUser(userEntity)
	if err != nil {
		return nil, err
	}

	err = userService.SendOtp(user.Username)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (userService *UserService) Login(userReq *model.UserLogin) (*model.UserLoginResponse, error) {
	user, err := userService.UserRepository.LoginUser(userReq)
	if err != nil {
		return nil, err
	}

	err = userService.Bcrypt.CompareAndHashPassword(user.Password, userReq.Password)
	if err != nil {
		return nil, err
	}

	token, err := userService.JWT.CreateToken(user.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := userService.JWT.CreateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	response := &model.UserLoginResponse{
		Username:     user.Username,
		Token:        token,
		RefreshToken: refreshToken,
	}
	return response, nil
}

func (userService *UserService) RefreshToken(token string) (string, error) {
	userId, err := userService.JWT.ValidateRefreshToken(token)
	if err != nil {
		return "", err
	}

	refreshToken, err := userService.JWT.CreateRefreshToken(userId)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (userService *UserService) GetUserById(id uuid.UUID) (*entity.User, error) {
	user, err := userService.UserRepository.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userService *UserService) GetUserByUsername(username string) (*entity.User, error) {
	user, err := userService.UserRepository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userService *UserService) UpdateUser(ctx *gin.Context, userReq *model.UpdateUser) (*model.UpdateUser, error) {
	userId, err := userService.JWT.GetLoginUser(ctx)
	if err != nil {
		return nil, err
	}

	if userReq.Password != "" {
		hashPassword, err := userService.Bcrypt.GenerateFromPassword(userReq.Password)
		if err != nil {
			return nil, err
		}
		userReq.Password = hashPassword
	}

	result, err := userService.UserRepository.UpdateUser(userId, userReq)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (userService *UserService) SendOtp(username string) error {
	user, otp, err := userService.UserRepository.SendOtp(username)
	if err != nil {
		return errors.New("failed to send otp to redis")
	}

	to := []string{user.Email}
	subject := "OTP Verification"
	message := fmt.Sprintf("Hello %s, your OTP is %s", user.Name, otp)

	err = userService.SMTP.SendMail(to, subject, message)

	if err != nil {
		return errors.New("failed to send otp email")
	}

	return nil
}

func (userService *UserService) ActivateUser(otpRequest *model.OtpRequest) error {
	err := userService.UserRepository.CheckOtp(otpRequest)
	if err != nil {
		return err
	}

	err = userService.UserRepository.ActivateUser(otpRequest.Username)
	if err != nil {
		return err
	}

	return nil
}

func (userService *UserService) GetIpAddress(ctx *gin.Context) string {
	ip := ctx.GetHeader("X-Forwarded-For")
	if ip == "" {
		ip = ctx.ClientIP()
	} else {
		ips := strings.Split(ip, ",")
		ip = strings.TrimSpace(ips[0])
	}
	return ip
}

func (userService *UserService) GetDeviceInfo(ctx *gin.Context) string {
	return ctx.GetHeader("User-Agent")
}
