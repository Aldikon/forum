package service

import (
	"context"

	"project/internal/repository"
	"project/model"
)

type UserService interface {
	Login(ctx context.Context, user *model.User) (context.Context, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewService(repository repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (u *userService) Login(ctx context.Context, user *model.User) (context.Context, error) {
	return nil, nil
}
