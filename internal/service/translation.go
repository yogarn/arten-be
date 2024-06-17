package service

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yogarn/arten/entity"
	"github.com/yogarn/arten/internal/repository"
	"github.com/yogarn/arten/pkg/deepl"
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

	if translation.OriginLanguage == "" && translation.TargetLanguage == "" && translation.Word == "" {
		return errors.New("missing required fields")
	}

	translation.CreatedAt = time.Now()
	translation.UpdatedAt = time.Now()

	translationResponse, err := deepl.Translate(translation.Word, translation.OriginLanguage, translation.TargetLanguage)
	if err != nil {
		return err
	}

	translation.Translate = translationResponse.Translations[0].Text

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
