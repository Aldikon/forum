package service

import (
	"forum/internal/model"
	service_commnet "forum/internal/service/service_social/comment"
	service_post "forum/internal/service/service_social/post"
	service_reaction "forum/internal/service/service_social/reaction"
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
