package repository

import (
	"database/sql"
)

type SesseansRepository interface{}

type sesseansRepository struct {
	db *sql.DB
}

func NewsesseansRepository(db *sql.DB) SesseansRepository {
	return &sesseansRepository{
		db: db,
	}
}

func (s *sesseansRepository) CreateSesseans() {
}
