package repository

import "database/sql"

type Repository struct {
	TranslationRepository ITranslationRepository
	UserRepository        IUserRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		TranslationRepository: NewTranslationRepository(db),
		UserRepository:        NewUserRepository(db),
	}
}
