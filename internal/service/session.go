package service

import (
	"github.com/google/uuid"
	"github.com/yogarn/arten/entity"
	"github.com/yogarn/arten/internal/repository"
	"github.com/yogarn/arten/pkg/jwt"
)

type ISessionService interface {
	CreateSession(session *entity.Session) error
	CheckSession(refreshToken string) error
	DeleteSession(userId uuid.UUID, refreshToken string) error
}

type SessionService struct {
	SessionRepository repository.ISessionRepository
	JWT               jwt.Interface
}

func NewSessionService(sessionRepository repository.ISessionRepository, jwt jwt.Interface) ISessionService {
	return &SessionService{
		SessionRepository: sessionRepository,
		JWT:               jwt,
	}
}

func (sessionService *SessionService) CreateSession(session *entity.Session) error {
	err := sessionService.SessionRepository.CreateSession(session)
	if err != nil {
		return err
	}
	return nil
}

func (sessionService *SessionService) CheckSession(refreshToken string) error {
	err := sessionService.SessionRepository.CheckSession(refreshToken)
	if err != nil {
		return err
	}
	return nil
}

func (sessionService *SessionService) DeleteSession(userId uuid.UUID, refreshToken string) error {
	err := sessionService.SessionRepository.DeleteSession(userId, refreshToken)
	if err != nil {
		return err
	}
	return nil
}
