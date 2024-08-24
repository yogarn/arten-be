package entity

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Id           uuid.UUID `json:"id" sql:"not null"`
	UserId       uuid.UUID `json:"userId" sql:"not null"`
	RefreshToken string    `json:"refreshToken" sql:"not null"`
	DeviceInfo   string    `json:"deviceInfo" sql:"not null"`
	IpAddress    string    `json:"ipAddress" sql:"not null"`
	CreatedAt    time.Time `json:"createdAt" sql:"not null; default:CURRENT_TIMESTAMP"`
}
