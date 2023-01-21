package repository

import "database/sql"

type CategoryRepository interface{}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *categoryRepository {
	return &categoryRepository{
		db: db,
	}
}
