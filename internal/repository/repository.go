package repository

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
)

type Repository struct {
	TranslationRepository ITranslationRepository
	UserRepository        IUserRepository
}

func NewRepository(db *sql.DB, redis *redis.Client) *Repository {
	return &Repository{
		TranslationRepository: NewTranslationRepository(db),
		UserRepository:        NewUserRepository(db, redis),
	}
}
