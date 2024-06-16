package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/yogarn/arten/entity"
)

type ITranslationRepository interface {
	CreateTranslation(translation *entity.Translation) error
	GetTranslation(id uuid.UUID) (*entity.Translation, error)
	UpdateTranslation(id uuid.UUID, translation *entity.Translation) error
	DeleteTranslation(id uuid.UUID) error
}

type TranslationRepository struct {
	db *sql.DB
}

func NewTranslationRepository(db *sql.DB) ITranslationRepository {
	return &TranslationRepository{db}
}

func (translationRepository *TranslationRepository) CreateTranslation(translation *entity.Translation) error {
	stmt := `
		INSERT INTO translations (id, origin_language, target_language, word, translate, updated_at, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	tx, err := translationRepository.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(stmt, translation.Id, translation.OriginLanguage, translation.TargetLanguage, translation.Word, translation.Translate, translation.UpdatedAt, translation.CreatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

func (translationRepository *TranslationRepository) GetTranslation(id uuid.UUID) (*entity.Translation, error) {
	stmt := `
		SELECT id, origin_language, target_language, word, translate, updated_at, created_at
		FROM translations
		WHERE id = $1
	`

	tx, err := translationRepository.db.Begin()
	if err != nil {
		return nil, err
	}
	row := tx.QueryRow(stmt, id)
	translation := &entity.Translation{}
	err = row.Scan(&translation.Id, &translation.OriginLanguage, &translation.TargetLanguage, &translation.Word, &translation.Translate, &translation.UpdatedAt, &translation.CreatedAt)
	if err != nil {
		return nil, err
	}
	return translation, nil
}

func (translationRepository *TranslationRepository) UpdateTranslation(id uuid.UUID, translation *entity.Translation) error {
	stmt := `
		UPDATE translations
		SET origin_language = $1, target_language = $2, word = $3, translate = $4, updated_at = $5
		WHERE id = $6
	`

	tx, err := translationRepository.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(stmt, translation.OriginLanguage, translation.TargetLanguage, translation.Word, translation.Translate, translation.UpdatedAt, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

func (translationRepository *TranslationRepository) DeleteTranslation(id uuid.UUID) error {
	stmt := `
		DELETE FROM translations
		WHERE id = $1
	`

	tx, err := translationRepository.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(stmt, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}
