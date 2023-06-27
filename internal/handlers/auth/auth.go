package handlers

import "forum/internal/model"

type auth struct {
	model.AuthService
}

func NewAuthorization(a model.AuthService) *auth {
	return &auth{
		AuthService: a,
	}
}
