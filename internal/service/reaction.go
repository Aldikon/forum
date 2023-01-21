package service

import "database/sql"

type ReactionService interface{}

type reactionService struct {
	db *sql.DB
}

func NewReactionService(db *sql.DB) *reactionService {
	return &reactionService{
		db: db,
	}
}
