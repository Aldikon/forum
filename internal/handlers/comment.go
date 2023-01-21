package handlers

import (
	"project/internal/service"
)

type CommentHandler interface{}

type commentHandler struct{}

func NewCommentHandler(userService service.UserService) *commentHandler {
	return &commentHandler{}
}
