package handlers

import "github.com/Aldikon/forum/internal/model"

type social struct {
	social model.SocialService
	auth   model.AuthService
}

func NewSocial(s model.SocialService, a model.AuthService) *social {
	return &social{
		social: s,
		auth:   a,
	}
}
