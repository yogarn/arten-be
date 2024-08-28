package service

import (
	"github.com/yogarn/arten/internal/repository"
	"github.com/yogarn/arten/pkg/bcrypt"
	"github.com/yogarn/arten/pkg/jwt"
	"github.com/yogarn/arten/pkg/smtp"
	"google.golang.org/grpc"
)

type Service struct {
	TranslationService ITranslationService
	UserService        IUserService
	SessionService     ISessionService
	TranscribeService  ITranscribeService
}

func NewService(repository *repository.Repository, bcrypt bcrypt.Interface, jwt jwt.Interface, smtp smtp.Interface, grpcConn *grpc.ClientConn) *Service {
	return &Service{
		TranslationService: NewTranslationService(repository.TranslationRepository, jwt),
		UserService:        NewUserService(repository.UserRepository, bcrypt, jwt, smtp),
		SessionService:     NewSessionService(repository.SessionRepository, jwt),
		TranscribeService:  NewTranscribeService(grpcConn),
	}
}
