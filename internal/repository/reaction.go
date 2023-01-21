package repository

import "database/sql"

type ReactionRepository interface{}

type reactionRepository struct {
	db *sql.DB
}

func NewReactionRepository(db *sql.DB) *reactionRepository {
	return &reactionRepository{
		db: db,
	}
}
