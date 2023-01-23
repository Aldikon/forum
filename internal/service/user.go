package service

import (
	"errors"

	"project/internal/dot"
	"project/internal/repository"
)

var (
	ErrIncorrectEmailInput    = errors.New("incorrect email input")
	ErrIncorrectNameInput     = errors.New("incorrect name input")
	ErrIncorrectPasswordInput = errors.New("incorrect password input")
	ErrNotTheSamePassword     = errors.New("not the same password")
)

type UserService interface {
	SignUp(*dot.UserSignUp) error
	LogIn(user *dot.UserLogIn) error
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (u *userService) SignUp(user *dot.UserSignUp) error {
	// if user.Password != user.ConfirmPassword {
	// 	return ErrNotTheSamePassword
	// }
	err := u.repository.CreateUser(user)
	return err
}

func (u *userService) LogIn(user *dot.UserLogIn) error {
	if err := u.repository.ReadToRegisterUser(user); err != nil {
		return err
	}

	return nil
}
