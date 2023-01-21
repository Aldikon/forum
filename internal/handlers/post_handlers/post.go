package handlers

import (
	"project/internal/service"
)

type PostHandler interface{}

type postHandler struct{}

func NewPostHandler(userService service.UserService) *postHandler {
	return &postHandler{}
}
