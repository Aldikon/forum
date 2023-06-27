package service

import (
	"forum/internal/model"
)

type auth struct {
	repo model.AuthorizationRepo
}

func NewAuth(r model.AuthorizationRepo) *auth {
	return &auth{
		repo: r,
	}
}
