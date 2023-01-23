package post_handlers

import (
	"net/http"

	"project/internal/service"
)

type PostHandler interface {
	CreatePost(w http.ResponseWriter, r *http.Request)
}

type postHandler struct {
	postService service.PostService
}

func NewPostHandler(postService service.PostService) PostHandler {
	return &postHandler{
		postService: postService,
	}
}
