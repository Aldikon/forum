package model

import "context"

type AuthorizationRepo interface {
	LogIn(ctx context.Context, user LogIn) (int64, error)
	AddUser(ctx context.Context, user Registration) error
	AddSession(ctx context.Context, session Session) error
	GetID(ctx context.Context, token string) (Session, error)
	GetByID(ctx context.Context, userID int64) (string, error)
	DeleteSession(ctx context.Context, userID int64) error
}

type PostRepo interface {
	Add(ctx context.Context, post CreatePost) error
	GetByID(ctx context.Context, postID int64) (Post, error)
	GetAll(ctx context.Context) ([]Post, error)
	GetCategoryAll(ctx context.Context) ([]string, error)
	GetByFilter(ctx context.Context, filter string) ([]Post, error)
	GetByLiked(ctx context.Context, userID int64) ([]Post, error)
}

type CommentRepo interface {
	Add(ctx context.Context, post CreateComment) error
	AddReply(ctx context.Context, post CreateComment) error
	GetByPostID(ctx context.Context, postID int64) ([]Comment, error)
}

type ReactionRepo interface {
	AddPost(ctx context.Context, reac CreateReactionPost) error
	AddComment(ctx context.Context, reac CreateReactionComment) error
}
