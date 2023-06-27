package service

import (
	"context"
	"time"

	"github.com/Aldikon/forum/internal/model"
)

type comment struct {
	repo     model.CommentRepo
	reaction model.ReactionService
}

func NewComment(r model.CommentRepo, rs model.ReactionService) *comment {
	return &comment{
		repo:     r,
		reaction: rs,
	}
}

func (c *comment) AddComment(ctx context.Context, comment model.CreateComment) error {
	var err error

	if err := comment.Validate(); err != nil {
		return err
	}

	comment.CreateAtt = time.Now()

	if comment.ParentID > 0 {
		err = c.addCommentReply(ctx, comment)
	} else {
		err = c.addComment(ctx, comment)
	}

	if err != nil {
		return err
	}
	return nil
}

func (c *comment) addComment(ctx context.Context, comment model.CreateComment) error {
	return c.repo.Add(ctx, comment)
}

func (c *comment) addCommentReply(ctx context.Context, comment model.CreateComment) error {
	return c.repo.AddReply(ctx, comment)
}

func (c *comment) GetCommentByPostID(ctx context.Context, postID int64) ([]model.Comment, error) {
	return c.repo.GetByPostID(ctx, postID)
}
