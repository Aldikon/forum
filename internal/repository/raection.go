package repository

import (
	"context"
	"database/sql"

	"forum/internal/model"
)

type reaction struct {
	db *sql.DB
}

func NewReactiona(db *sql.DB) *reaction {
	return &reaction{
		db: db,
	}
}

func (r *reaction) AddPost(ctx context.Context, reac model.CreateReactionPost) error {
	query := `INSERT OR REPLACE INTO Reactions_Posts (user_id,post_id,type) 
	VALUES ($1 , $2 , 
		IIF($3 = (SELECT type FROM Reactions_Posts WHERE post_id = $2 AND user_id = $1), 0, $3))`

	res, err := r.db.ExecContext(ctx, query, reac.UserID, reac.PostID, reac.Type)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count <= 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *reaction) AddComment(ctx context.Context, reac model.CreateReactionComment) error {
	query := `INSERT OR REPLACE INTO Reactions_Comments (user_id,comment_id,type) 
	VALUES ($1 , $2 , 
		IIF($3 = (SELECT type FROM Reactions_Comments WHERE comment_id = $2 AND user_id = $1), 0, $3))`

	res, err := r.db.ExecContext(ctx, query, reac.UserID, reac.CommentID, reac.Type)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count <= 0 {
		return sql.ErrNoRows
	}
	return nil
}
