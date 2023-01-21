package service

import "database/sql"

type CommentService interface{}

type commentService struct {
	db *sql.DB
}

func NewCommentService(db *sql.DB) *commentService {
	return &commentService{
		db: db,
	}
}
