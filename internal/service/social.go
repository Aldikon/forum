package service

import (
	"github.com/Aldikon/forum/internal/model"
	service_commnet "github.com/Aldikon/forum/internal/service/service_social/comment"
	service_post "github.com/Aldikon/forum/internal/service/service_social/post"
	service_reaction "github.com/Aldikon/forum/internal/service/service_social/reaction"
)

type social struct {
	model.PostService
	model.CommentService
	model.ReactionService
}

func NewSocial(pr model.PostRepo, cr model.CommentRepo, rr model.ReactionRepo) *social {
	rs := service_reaction.NewReaction(rr)
	return &social{
		PostService:     service_post.NewPost(pr, rs),
		CommentService:  service_commnet.NewComment(cr, rs),
		ReactionService: rs,
	}
}
