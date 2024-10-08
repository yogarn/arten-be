package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/yogarn/arten/entity"
)

type ITranslationRepository interface {
	CreateTranslation(translation *entity.Translation) error
	GetTranslation(id uuid.UUID) (*entity.Translation, error)
	GetTranslationHistory(userId uuid.UUID) ([]entity.Translation, error)
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
		INSERT INTO translations (id, user_id, origin_language, target_language, word, translate, created_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	tx, err := translationRepository.db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %v", err)
	}

	_, err = tx.Exec(stmt, translation.Id, translation.UserId, translation.OriginLanguage, translation.TargetLanguage, translation.Word, translation.Translate, translation.CreatedAt)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("exec statement: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit transaction: %v", err)
	}

	return nil
}

func (translationRepository *TranslationRepository) GetTranslation(id uuid.UUID) (*entity.Translation, error) {
	stmt := `
		SELECT id, user_id, origin_language, target_language, word, translate, created_at
		FROM translations
		WHERE id = ?
	`

	row := translationRepository.db.QueryRow(stmt, id)
	translation := &entity.Translation{}
	err := row.Scan(&translation.Id, &translation.UserId, &translation.OriginLanguage, &translation.TargetLanguage, &translation.Word, &translation.Translate, &translation.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no record found")
		}
		return nil, fmt.Errorf("scan row: %v", err)
	}

	return translation, nil
}

func (translationRepository *TranslationRepository) GetTranslationHistory(userId uuid.UUID) ([]entity.Translation, error) {
	stmt := `
		SELECT id, user_id, origin_language, target_language, word, translate, created_at
		FROM translations
		WHERE user_id = ?
		LIMIT 10
	`

	rows, err := translationRepository.db.Query(stmt, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no record found")
		}
		return nil, fmt.Errorf("scan row: %v", err)
	}

	defer rows.Close()

	translations := []entity.Translation{}
	for rows.Next() {
		translation := entity.Translation{}
		err := rows.Scan(&translation.Id, &translation.UserId, &translation.OriginLanguage, &translation.TargetLanguage, &translation.Word, &translation.Translate, &translation.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("scan row: %v", err)
		}
		translations = append(translations, translation)
	}

	return translations, nil
}

func (translationRepository *TranslationRepository) UpdateTranslation(id uuid.UUID, translation *entity.Translation) error {
	stmt := `
		UPDATE translations
		SET origin_language = ?, target_language = ?, word = ?, translate = ?
		WHERE id = ?
	`

	tx, err := translationRepository.db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %v", err)
	}

	_, err = tx.Exec(stmt, translation.OriginLanguage, translation.TargetLanguage, translation.Word, translation.Translate, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("exec statement: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit transaction: %v", err)
	}

	return nil
}

func (translationRepository *TranslationRepository) DeleteTranslation(id uuid.UUID) error {
	stmt := `
		DELETE FROM translations
		WHERE id = ?
	`

	tx, err := translationRepository.db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %v", err)
	}

	_, err = tx.Exec(stmt, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("exec statement: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit transaction: %v", err)
	}

	return nil
}
