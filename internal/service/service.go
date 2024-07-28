package service

import (
	"github.com/yogarn/arten/internal/repository"
	"github.com/yogarn/arten/pkg/bcrypt"
	"github.com/yogarn/arten/pkg/jwt"
)

type Service struct {
	TranslationService ITranslationService
	UserService        IUserService
}

func NewService(repository *repository.Repository, bcrypt bcrypt.Interface, jwt jwt.Interface) *Service {
	return &Service{
		TranslationService: NewTranslationService(repository.TranslationRepository),
		UserService:        NewUserService(repository.UserRepository, bcrypt, jwt),
	}
}
