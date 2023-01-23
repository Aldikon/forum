package post_handlers

import (
	"project/internal/service"
)

type PostHandler interface{}

type postHandler struct {
	postService *service.PostService
}

func NewPostHandler(postService *service.PostService) *postHandler {
	return &postHandler{
		postService: postService,
	}
}
