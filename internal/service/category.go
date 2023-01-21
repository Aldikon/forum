package service

import "database/sql"

type CategoryService interface{}

type categoryService struct {
	db *sql.DB
}

func NewCategoryService(db *sql.DB) *categoryService {
	return &categoryService{
		db: db,
	}
}
