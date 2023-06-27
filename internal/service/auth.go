package service

import (
	"github.com/Aldikon/forum/internal/model"
	service "github.com/Aldikon/forum/internal/service/service_auth"
)

type auth struct {
	model.AuthService
}

func NewAuthorization(r model.AuthorizationRepo) *auth {
	return &auth{
		service.NewAuth(r),
	}
}
