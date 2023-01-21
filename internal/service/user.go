package service

import (
	"context"
	"errors"

	"project/internal/dot"
	"project/internal/repository"
	"project/model"
)

var (
	ErrIncorrectEmailInput    = errors.New("incorrect email input")
	ErrIncorrectNameInput     = errors.New("incorrect name input")
	ErrIncorrectPasswordInput = errors.New("incorrect password input")
	ErrNotTheSamePassword     = errors.New("not the same password")
)

type UserService interface {
	SignUp(*dot.UserSignUp) error
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

func (u *userService) SignUp(user *dot.UserSignUp) error {
	if user.Password != user.ConfirmPassword {
		return ErrNotTheSamePassword
	}
	return nil
}

func (u *userService) Login(ctx context.Context, user *model.User) (context.Context, error) {
	return nil, nil
}
