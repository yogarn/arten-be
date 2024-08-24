package service

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yogarn/arten/entity"
	"github.com/yogarn/arten/internal/repository"
	"github.com/yogarn/arten/pkg/deepl"
	"github.com/yogarn/arten/pkg/jwt"
)

type ITranslationService interface {
	CreateTranslation(ctx *gin.Context, translation *entity.Translation) error
	GetTranslation(ctx *gin.Context, id uuid.UUID) (*entity.Translation, error)
	GetTranslationHistory(ctx *gin.Context) ([]entity.Translation, error)
	UpdateTranslation(ctx *gin.Context, id uuid.UUID, translation *entity.Translation) error
	DeleteTranslation(ctx *gin.Context, id uuid.UUID) error
}

type TranslationService struct {
	TranslationRepository repository.ITranslationRepository
	JWT                   jwt.Interface
}

func NewTranslationService(translationRepository repository.ITranslationRepository, jwt jwt.Interface) ITranslationService {
	return &TranslationService{
		TranslationRepository: translationRepository,
		JWT:                   jwt,
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

	userId, err := translationService.JWT.GetLoginUser(ctx)
	if err != nil {
		return err
	}

	translation.UserId = userId

	translationResponse, err := deepl.Translate(translation.Word, translation.OriginLanguage, translation.TargetLanguage)
	if err != nil {
		return err
	}

	translation.Translate = translationResponse.Translations[0].Text

	return translationService.TranslationRepository.CreateTranslation(translation)
}

func (translationService *TranslationService) GetTranslation(ctx *gin.Context, id uuid.UUID) (*entity.Translation, error) {
	translation, err := translationService.TranslationRepository.GetTranslation(id)
	if err != nil {
		return nil, err
	}
	return translation, nil
}

func (translationService *TranslationService) GetTranslationHistory(ctx *gin.Context) ([]entity.Translation, error) {
	userId, err := translationService.JWT.GetLoginUser(ctx)
	if err != nil {
		return nil, err
	}

	translations, err := translationService.TranslationRepository.GetTranslationHistory(userId)
	if err != nil {
		return nil, err
	}
	return translations, nil
}

// NOTE: should be removed in near future, not needed
func (translationService *TranslationService) UpdateTranslation(ctx *gin.Context, id uuid.UUID, translation *entity.Translation) error {
	err := translationService.TranslationRepository.UpdateTranslation(id, translation)
	if err != nil {
		return err
	}
	return nil
}

func (translationService *TranslationService) DeleteTranslation(ctx *gin.Context, id uuid.UUID) error {
	err := translationService.TranslationRepository.DeleteTranslation(id)
	if err != nil {
		return err
	}
	return nil
}
