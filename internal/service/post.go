package service

import "database/sql"

type PostService interface{}

type postService struct {
	db *sql.DB
}

func NewPostService(db *sql.DB) *postService {
	return &postService{
		db: db,
	}
}
