package service

import "github.com/yogarn/arten/internal/repository"

type Service struct {
	TranslationService ITranslationService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		TranslationService: NewTranslationService(repository.TranslationRepository),
	}
}
