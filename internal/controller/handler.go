package controller

import (
	"net/http"

	"project/internal/controller/controllers"
	"project/internal/service"
)

type ErrorInterface interface {
	Error(http.ResponseWriter, *http.Request)
}

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
	ErrorInterface
	MiddlewareInterface
	SignInInterface
	SignUpInterface
}

type controller struct {
	ErrorInterface
	MiddlewareInterface
	SignInInterface
	SignUpInterface
}

func NewController(s service.Service) *controller {
	return &controller{
		ErrorInterface:      controllers.NewErrorController(s),
		MiddlewareInterface: controllers.NewMiddlewareController(),
		SignInInterface:     controllers.NewSingInController(s),
		SignUpInterface:     controllers.NewSingUpController(s),
	}
}
