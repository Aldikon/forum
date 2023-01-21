package user_handlers

import (
	"net/http"

	"project/internal/service"
)

type UserHandler interface {
	LogIn(http.ResponseWriter, *http.Request)
	LogOut(http.ResponseWriter, *http.Request)
	Profile(http.ResponseWriter, *http.Request)
	SignUp(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}
