package model

import "context"

type AuthService interface {
	Add(ctx context.Context, user Registration) error
	GetID(ctx context.Context, token string) (int64, error)
	GetByID(ctx context.Context, userID int64) (string, error)
	LogIn(ctx context.Context, user LogIn) (Session, error)
	LogOut(ctx context.Context, userID int64) error
}

type PostService interface {
	AddPost(ctx context.Context, post CreatePost) error
	GetPostByID(ctx context.Context, postID int64) (Post, error)
	GetPostAll(ctx context.Context) ([]Post, error)
	GetPostAllFilter(ctx context.Context, userID int64, filter string) ([]Post, error)
	GetCategoryAll(ctx context.Context) ([]string, error)
}

type CommentService interface {
	AddComment(ctx context.Context, post CreateComment) error
	GetCommentByPostID(ctx context.Context, postID int64) ([]Comment, error)
}

type ReactionService interface {
	AddReactionPost(ctx context.Context, reac CreateReactionPost) error
	AddReactionComment(ctx context.Context, reac CreateReactionComment) error
}

type SocialService interface {
	PostService
	CommentService
	ReactionService
}
