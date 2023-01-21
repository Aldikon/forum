package repository

import (
	"context"
	"database/sql"

	"project/model"
)

type UserRepository interface{}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUser(ctx context.Context, user *model.User) (context.Context, error) {
	return nil, nil
}
