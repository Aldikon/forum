package service

import (
	"context"

	"forum/internal/model"
)

type reaction struct {
	repo model.ReactionRepo
}

func NewReaction(r model.ReactionRepo) *reaction {
	return &reaction{
		repo: r,
	}
}

func (r *reaction) AddReactionPost(ctx context.Context, reac model.CreateReactionPost) error {
	return r.repo.AddPost(ctx, reac)
}

func (r *reaction) AddReactionComment(ctx context.Context, reac model.CreateReactionComment) error {
	return r.repo.AddComment(ctx, reac)
}
