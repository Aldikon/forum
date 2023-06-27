package service

import (
	"context"
	"errors"
	"time"

	"github.com/Aldikon/forum/internal/model"
)

type post struct {
	repo     model.PostRepo
	reaction model.ReactionService
}

func NewPost(r model.PostRepo, rs model.ReactionService) *post {
	return &post{
		repo:     r,
		reaction: rs,
	}
}

func (p *post) AddPost(ctx context.Context, post model.CreatePost) error {
	if err := post.Validate(); err != nil {
		return err
	}

	post.CreateAtt = time.Now()

	if err := p.repo.Add(ctx, post); err != nil {
		return err
	}

	return nil
}

func (p *post) GetPostByID(ctx context.Context, postID int64) (model.Post, error) {
	var post model.Post
	post, err := p.repo.GetByID(ctx, postID)
	if err != nil {
		return post, err
	}

	return post, err
}

func (p *post) GetPostAll(ctx context.Context) ([]model.Post, error) {
	return p.repo.GetAll(ctx)
}

func (p *post) GetPostAllFilter(ctx context.Context, userID int64, filter string) ([]model.Post, error) {
	if filter == "liked" {
		if userID <= 0 {
			return nil, errors.New("...")
		}
		return p.repo.GetByLiked(ctx, userID)
	}

	return p.repo.GetByFilter(ctx, filter)
}

func (p *post) GetCategoryAll(ctx context.Context) ([]string, error) {
	return p.repo.GetCategoryAll(ctx)
}
