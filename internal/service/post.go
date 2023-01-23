package service

import (
	"project/internal/dot"
	"project/internal/repository"
)

type PostService interface {
	CreatePost(*dot.CreatePost) error
}

type postService struct {
	postRepository repository.PostRepository
	userRepository repository.UserRepository
}

func NewPostService(postRepo repository.PostRepository, uerRepo repository.UserRepository) PostService {
	return &postService{
		postRepository: postRepo,
		userRepository: uerRepo,
	}
}

func (p *postService) CreatePost(post *dot.CreatePost) error {
	_, err := p.userRepository.ReadToIdUser(post.UserId)
	if err != nil {
		return err
	}
	if err := p.postRepository.CreatePost(post); err != nil {
		return err
	}
	return nil
}
