package service

import (
	"project/internal/model"
	"project/internal/repository"
	"project/internal/service/services"
)

type UserService interface {
	SignInService(model.User) error
	SignUpService(model.User, string) error
}

type Service interface {
	UserService
}

type service struct {
	UserService
}

func NewService(repo repository.Repository) *service {
	return &service{
		UserService: services.NewUserService(repo),
	}
}
