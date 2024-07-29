package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yogarn/arten/entity"
	"github.com/yogarn/arten/internal/repository"
	"github.com/yogarn/arten/model"
	"github.com/yogarn/arten/pkg/bcrypt"
	"github.com/yogarn/arten/pkg/jwt"
)

type IUserService interface {
	Register(userReq *model.UserRegister) (*entity.User, error)
	Login(userReq *model.UserLogin) (*model.UserLoginResponse, error)
	GetUserById(id uuid.UUID) (*entity.User, error)
	UpdateUser(ctx *gin.Context, user *model.UpdateUser) (*model.UpdateUser, error)
	SendOtp(username string) error
	ActivateUser(otpRequest *model.OtpRequest) error
}

type UserService struct {
	UserRepository       repository.IUserRepository
	Bcrypt               bcrypt.Interface
	JWT                  jwt.Interface
	ProfilePictureBucket string
}

func NewUserService(userRepository repository.IUserRepository, bcrypt bcrypt.Interface, jwt jwt.Interface) IUserService {
	return &UserService{
		UserRepository:       userRepository,
		Bcrypt:               bcrypt,
		JWT:                  jwt,
		ProfilePictureBucket: "profilePicture",
	}
}

func (userService *UserService) Register(userReq *model.UserRegister) (*entity.User, error) {
	hashPassword, err := userService.Bcrypt.GenerateFromPassword(userReq.Password)
	if err != nil {
		return nil, err
	}
	userEntity := &entity.User{
		ID:       uuid.New(),
		Username: userReq.Username,
		Password: hashPassword,
		Name:     userReq.Name,
		Email:    userReq.Email,
	}
	user, err := userService.UserRepository.CreateUser(userEntity)
	if err != nil {
		return nil, err
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

	response := &model.UserLoginResponse{
		Username: user.Username,
		Token:    token,
	}
	return response, nil
}

func (userService *UserService) GetUserById(id uuid.UUID) (*entity.User, error) {
	user, err := userService.UserRepository.GetUserById(id)
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
	err := userService.UserRepository.SendOtp(username)
	if err != nil {
		return err
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
