package service

import (
	"forum/internal/model"
	service "forum/internal/service/service_auth"
)

type auth struct {
	model.AuthService
}

func NewAuthorization(r model.AuthorizationRepo) *auth {
	return &auth{
		service.NewAuth(r),
	}
}
