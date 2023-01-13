package services

import (
	"errors"
	"net/http"

	"project/internal/model"
	"project/internal/repository"
)

type userService struct {
	repository repository.UserRepository
}

func NewUserService(r repository.UserRepository) *userService {
	return &userService{
		repository: r,
	}
}

func (u *userService) SignInService(user model.User) error {
	if user.Email == "" || user.Password == "" {
		return errors.New("Emty login or password!")
	}

	if err := u.repository.ReadToRegisterUser(user); err != nil {
		return err
	}
	return nil
}

func (u *userService) SignUpService(user model.User, confirmPassword string) error {
	if err := user.CheckInput(confirmPassword); err != nil {
		return err
	}

	if err := u.repository.CreateUser(user); err != nil {
		return errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return nil
}
