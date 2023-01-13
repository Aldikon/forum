// bla bla
package repository

// ОПИСАНИЕ СЛОЯ РЕПОЗИТОРИЯ

import (
	"database/sql"

	"project/internal/model"
	"project/internal/repository/repositories"
)

type UserRepository interface {
	CreateUser(model.User) error
	ReadToRegisterUser(user model.User) error
	ReadUser(string, string) (model.User, error)
	UpdateUser(model.User, string) error
	DeleteUser(model.User) error
}

type PostRepository interface {
	// CreatePost(model.Post) error
	// ReadAllPost() ([]model.Post, error)
	// ReadByUserPost(int) ([]model.Post, error)
	// ReadByIDPost(int) (model.Post, error)
	// UpdatePost(model.Post, int) error
	// DeletePost(int) error
}

type CommentRepository interface {
	CreateComment(model.Comment) error
	ReadByPostIDComment(int) ([]model.Comment, error)
	ReadByIDComment(int) (model.Comment, error)
	UpdateComment(model.Comment) error
	DeleteComment(int) error
}

type CategoriRepository interface{}

type CommentReactionRepository interface {
	CreateTheCommentReaction(int, int, int) error
	UpdateTheCommentReaction(int, int) error
	ReadAllCommentReaction(int) ([]int, error)
	DeleteTheCommentReaction(int) error
}

type PostReactionRepository interface {
	CreateThePostReaction(int, int, int) error
	UpdateThePostReaction(int, int) error
	ReadAllPostReaction(int) ([]int, error)
	DeleteThePostReaction(int) error
}

type ReactionRepository interface {
	CommentReactionRepository
	PostReactionRepository
}

type Repository interface {
	UserRepository
	PostRepository
	CommentRepository
	CategoriRepository
	ReactionRepository
}

type repository struct {
	UserRepository
	PostRepository
	CommentRepository
	CategoriRepository
	ReactionRepository
}

type reaction struct {
	CommentReactionRepository
	PostReactionRepository
}

// *repository
// Создает репозитори
func NewRepository(db *sql.DB) *repository {
	return &repository{
		UserRepository:     repositories.NewUserRepository(db),
		PostRepository:     repositories.NewPostRepository(db),
		CommentRepository:  repositories.NewCommentRepository(db),
		CategoriRepository: repositories.NewCategoriRepository(db),
		ReactionRepository: newReactionRepository(db),
	}
}

func newReactionRepository(db *sql.DB) *reaction {
	return &reaction{
		CommentReactionRepository: repositories.NewReactionCommentRepository(db),
		PostReactionRepository:    repositories.NewPostReactionRepository(db),
	}
}
