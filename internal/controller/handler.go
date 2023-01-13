package controller

import (
	"net/http"

	"project/internal/controller/controllers"
	"project/internal/service"
)

type MiddlewareInterface interface {
	SignUpMiddleware(http.Handler) http.Handler
}

type SignInInterface interface {
	SignIn(w http.ResponseWriter, r *http.Request)
}

type SignUpInterface interface {
	SignUp(w http.ResponseWriter, r *http.Request)
}

type Controller interface {
	MiddlewareInterface
	SignInInterface
	SignUpInterface
}

type controller struct {
	MiddlewareInterface
	SignInInterface
	SignUpInterface
}

func NewController(s service.Service) *controller {
	return &controller{
		MiddlewareInterface: controllers.NewMiddlewareController(),
		SignInInterface:     controllers.NewSingInController(s),
		SignUpInterface:     controllers.NewSingUpController(s),
	}
}
