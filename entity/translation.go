package entity

import (
	"time"

	"github.com/google/uuid"
)

type Translation struct {
	Id             uuid.UUID `json:"id" sql:"not null"`
	ChatId         uuid.UUID `json:"chatId" sql:"not null"`
	OriginLanguage string    `json:"originLanguage" sql:"not null"`
	TargetLanguage string    `json:"targetLanguage" sql:"not null"`
	Word           string    `json:"word" sql:"not null"`
	Translate      string    `json:"translate" sql:"not null"`
	UpdatedAt      time.Time `json:"updatedAt" sql:"not null"`
	CreatedAt      time.Time `json:"createdAt" sql:"not null; default:CURRENT_TIMESTAMP"`
}
