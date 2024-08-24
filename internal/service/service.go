package service

import (
	"github.com/yogarn/arten/internal/repository"
	"github.com/yogarn/arten/pkg/bcrypt"
	"github.com/yogarn/arten/pkg/jwt"
	"github.com/yogarn/arten/pkg/smtp"
)

type Service struct {
	TranslationService ITranslationService
	UserService        IUserService
	SessionService     ISessionService
}

func NewService(repository *repository.Repository, bcrypt bcrypt.Interface, jwt jwt.Interface, smtp smtp.Interface) *Service {
	return &Service{
		TranslationService: NewTranslationService(repository.TranslationRepository, jwt),
		UserService:        NewUserService(repository.UserRepository, bcrypt, jwt, smtp),
		SessionService:     NewSessionService(repository.SessionRepository, jwt),
	}
}
