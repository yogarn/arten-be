package repository

import "database/sql"

type Repository struct {
	TranslationRepository ITranslationRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		TranslationRepository: NewTranslationRepository(db),
	}
}
