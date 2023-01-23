package service

import (
	"project/internal/repository"
)

type PostService interface{}

type postService struct {
	PostRepository *repository.PostRepository
}

func NewPostService(postRepo *repository.PostRepository) *postService {
	return &postService{
		PostRepository: postRepo,
	}
}

func (p *postService) CreatePost() error {
	return nil
}
