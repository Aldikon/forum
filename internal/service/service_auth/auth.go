package service

import "github.com/Aldikon/forum/internal/model"

type auth struct {
	repo model.AuthorizationRepo
}

func NewAuth(r model.AuthorizationRepo) *auth {
	return &auth{
		repo: r,
	}
}
