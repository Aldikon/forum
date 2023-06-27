package repository

import (
	"context"
	"database/sql"
	"time"

	"forum/internal/model"
)

type comment struct {
	db *sql.DB
}

func NewComment(db *sql.DB) *comment {
	return &comment{
		db: db,
	}
}

func (c *comment) Add(ctx context.Context, commment model.CreateComment) error {
	query := `INSERT INTO Comments (user_id, post_id, create_att, content) 
	VALUES ($1, $2, $3, $4);`
	res, err := c.db.ExecContext(ctx, query, commment.UserID, commment.PostID, commment.CreateAtt, commment.Content)
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

func (c *comment) AddReply(ctx context.Context, commment model.CreateComment) error {
	query := `INSERT INTO Comments (user_id,post_id,parent_id,content) 
	VALUES ($1, $2, $3, $4);`
	res, err := c.db.ExecContext(ctx, query, commment.UserID, commment.PostID, commment.ParentID, commment.Content)
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

func (c *comment) GetByPostID(ctx context.Context, postID int64) ([]model.Comment, error) {
	comments := make([]model.Comment, 0)
	query := `SELECT Comments.id, Users.name, Comments.content, Comments.parent_id, Comments.create_att,
		COALESCE(SUM(CASE WHEN Reactions_Comments.type = 1 THEN 1 ELSE 0 END), 0) AS likes,
		COALESCE(SUM(CASE WHEN Reactions_Comments.type = -1 THEN 1 ELSE 0 END), 0) AS dislikes
	FROM Comments
	JOIN Users ON Comments.user_id = Users.id
	LEFT JOIN Reactions_Comments ON Comments.id = Reactions_Comments.comment_id
	WHERE Comments.post_id = $1
	GROUP BY Comments.id, Users.name;`

	// 	SELECT Comments.id, Users.name, Comments.content, Comments.parent_id,
	//        COALESCE(SUM(CASE WHEN Reactions_Comments.type = 1 THEN 1 ELSE 0 END), 0) AS likes,
	//        COALESCE(SUM(CASE WHEN Reactions_Comments.type = -1 THEN 1 ELSE 0 END), 0) AS dislikes
	// FROM Comments
	// JOIN Users ON Comments.user_id = Users.id
	// LEFT JOIN Reactions_Comments ON Comments.id = Reactions_Comments.comment_id
	// GROUP BY Comments.id, Users.name;

	rows, err := c.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c := model.Comment{}
		p := sql.NullInt64{}
		var createAttStr string
		err := rows.Scan(&c.ID, &c.UserName, &c.Content, &p, &createAttStr, &c.Like, &c.DisLike)
		if err != nil {
			return nil, err
		}
		c.CreateAtt, err = time.Parse("2006-01-02 15:04:05.999999999-07:00", createAttStr)
		if err != nil {
			return nil, err
		}
		if p.Valid {
			c.ParentID = p.Int64
		}

		comments = append(comments, c)
	}

	return comments, nil
}
