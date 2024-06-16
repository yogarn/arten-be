package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yogarn/arten/entity"
	"github.com/yogarn/arten/internal/repository"
)

type ITranslationService interface {
	CreateTranslation(ctx *gin.Context, translation *entity.Translation) error
	GetTranslation(ctx *gin.Context, id uuid.UUID) (*entity.Translation, error)
	UpdateTranslation(ctx *gin.Context, id uuid.UUID, translation *entity.Translation) error
	DeleteTranslation(ctx *gin.Context, id uuid.UUID) error
}

type TranslationService struct {
	TranslationRepository repository.ITranslationRepository
}

func NewTranslationService(translationRepository repository.ITranslationRepository) ITranslationService {
	return &TranslationService{
		TranslationRepository: translationRepository,
	}
}

func (translationService *TranslationService) CreateTranslation(ctx *gin.Context, translation *entity.Translation) error {
	if translation.Id == uuid.Nil {
		translation.Id = uuid.New()
	}
	return translationService.TranslationRepository.CreateTranslation(translation)
}

func (translationService *TranslationService) GetTranslation(ctx *gin.Context, id uuid.UUID) (*entity.Translation, error) {
	return translationService.TranslationRepository.GetTranslation(id)
}

func (translationService *TranslationService) UpdateTranslation(ctx *gin.Context, id uuid.UUID, translation *entity.Translation) error {
	return translationService.TranslationRepository.UpdateTranslation(id, translation)
}

func (translationService *TranslationService) DeleteTranslation(ctx *gin.Context, id uuid.UUID) error {
	return translationService.TranslationRepository.DeleteTranslation(id)
}
