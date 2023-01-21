package handlers

import (
	"project/internal/service"
)

type CategoryHandler interface{}

type categoryHandler struct{}

func NewCategoryHandler(userService service.UserService) *categoryHandler {
	return &categoryHandler{}
}
