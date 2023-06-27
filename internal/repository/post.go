package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"forum/internal/model"

	"github.com/mattn/go-sqlite3"
)

type post struct {
	db *sql.DB
}

func NewPost(db *sql.DB) *post {
	return &post{
		db: db,
	}
}

func (p *post) Add(ctx context.Context, post model.CreatePost) error {
	tx, err := p.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	queryPost := `INSERT INTO Posts (user_id,create_att,title,content)
	VALUES ($1, $2, $3, $4);`

	res, err := tx.ExecContext(ctx, queryPost, post.UserID, post.CreateAtt, post.Title, post.Content)
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

	postID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	queryCategory := `INSERT INTO Categories_Posts (post_id,category_id)
	VALUES ($1, (SELECT id FROM Categories WHERE Categories.name = $2));`

	for _, category := range post.Categories {

		res, err := tx.ExecContext(ctx, queryCategory, postID, category)
		if err != nil {
			var sqliteErr sqlite3.Error
			if errors.As(err, &sqliteErr) {
				if errors.Is(err, sqlite3.ErrConstraintNotNull) {
					return model.ErrCategoryExist
				}
			}
			return err
		}

		count, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if count <= 0 {
			return sql.ErrNoRows
		}

	}

	tx.Commit()
	return nil
}

func (p *post) GetByID(ctx context.Context, postID int64) (model.Post, error) {
	post := model.Post{}

	tx, err := p.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return post, err
	}

	queryPost := `SELECT Posts.id, Users.name, Posts.title, Posts.content, Posts.create_att,
		COALESCE(SUM(CASE WHEN Reactions_Posts.type = 1 THEN 1 ELSE 0 END), 0) as 'likes',	
	    COALESCE(SUM(CASE WHEN Reactions_Posts.type = -1 THEN 1 ELSE 0 END), 0) as 'dislikes'
	FROM Posts
	JOIN Users on Posts.user_id = Users.id
	LEFT JOIN Reactions_Posts on Posts.id = Reactions_Posts.post_id
	WHERE Posts.id = $1;`

	row := tx.QueryRowContext(ctx, queryPost, postID)

	var createAttStr string

	err = row.Scan(&post.ID, &post.UserName, &post.Title, &post.Content, &createAttStr, &post.Like, &post.DisLike)
	if err != nil {
		return post, err
	}
	post.CreateAtt, err = time.Parse("2006-01-02 15:04:05.999999999-07:00", createAttStr)
	if err != nil {
		return post, err
	}

	queryCategory := `
	SELECT  Categories.name
	FROM Categories_Posts 
	JOIN Categories
		ON Categories_Posts.category_id = Categories.id
	WHERE post_id = ?;`

	rows, err := tx.QueryContext(ctx, queryCategory, postID)
	if err != nil {
		return post, err
	}

	for rows.Next() {
		var category string
		err := rows.Scan(&category)
		if err != nil {
			return post, err
		}
		post.Categories = append(post.Categories, category)
	}

	return post, nil
}

func (p *post) GetAll(ctx context.Context) ([]model.Post, error) {
	posts := make([]model.Post, 0)

	tx, err := p.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return nil, err
	}

	query := `SELECT Posts.id, Users.name, Posts.create_att, Posts.title, Posts.content,
	(select count(*) FROM Reactions_Posts where post_id = Posts.id and type = 1) AS likes,
	(select count(*) FROM Reactions_Posts where post_id = Posts.id and type =  -1)  AS dislikes,
	(SELECT COUNT(*) FROM Comments WHERE Comments.post_id = Posts.id) AS sum_comments
FROM Posts
JOIN Users ON Posts.user_id = Users.id
ORDER by Posts.create_att DESC;`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	queryCategory := `
	SELECT  Categories.name
	FROM Categories_Posts 
	JOIN Categories
		ON Categories_Posts.category_id = Categories.id
	WHERE post_id = ?;`

	for rows.Next() {
		p := model.Post{}
		var timeStr string
		err := rows.Scan(&p.ID, &p.UserName, &timeStr, &p.Title, &p.Content, &p.Like, &p.DisLike, &p.SumComments)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		p.CreateAtt, err = time.Parse("2006-01-02 15:04:05.999999999-07:00", timeStr)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		rows, err := tx.QueryContext(ctx, queryCategory, p.ID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		for rows.Next() {
			var category string

			err := rows.Scan(&category)
			if err != nil {
				tx.Rollback()
				return nil, err
			}

			p.Categories = append(p.Categories, category)
		}

		posts = append(posts, p)
	}

	return posts, nil
}

func (p *post) GetCategoryAll(ctx context.Context) ([]string, error) {
	query := `SELECT Categories.name FROM Categories;`

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	categories := make([]string, 0)
	for rows.Next() {
		var category string
		err := rows.Scan(&category)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (p *post) GetByFilter(ctx context.Context, filter string) ([]model.Post, error) {
	posts := make([]model.Post, 0)

	tx, err := p.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return nil, err
	}

	query := `
	SELECT Posts.id, Users.name, Posts.create_att, Posts.title, Posts.content,
    	(select count(*) FROM Reactions_Posts where post_id = Posts.id and type = 1) AS likes,
    	(select count(*) FROM Reactions_Posts where post_id = Posts.id and type =  -1)  AS dislikes,
    	(SELECT COUNT(*) FROM Comments WHERE Comments.post_id = Posts.id) AS sum_comments
    FROM Posts
	JOIN Users ON Posts.user_id = Users.id
    JOIN Categories_Posts on Categories_Posts.post_id = Posts.id
    JOIN Categories on Categories.id = Categories_Posts.category_id
    WHERE Categories.name = $1
	ORDER by Posts.create_att DESC;`

	rows, err := tx.QueryContext(ctx, query, filter)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	queryCategory := `
	SELECT  Categories.name
	FROM Categories_Posts 
	JOIN Categories
		ON Categories_Posts.category_id = Categories.id
	WHERE post_id = ?;`

	for rows.Next() {
		p := model.Post{}
		var timeStr string
		err := rows.Scan(&p.ID, &p.UserName, &timeStr, &p.Title, &p.Content, &p.Like, &p.DisLike, &p.SumComments)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		p.CreateAtt, err = time.Parse("2006-01-02 15:04:05.999999999-07:00", timeStr)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		rows, err := tx.QueryContext(ctx, queryCategory, p.ID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		for rows.Next() {
			var category string

			err := rows.Scan(&category)
			if err != nil {
				tx.Rollback()
				return nil, err
			}

			p.Categories = append(p.Categories, category)
		}

		posts = append(posts, p)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *post) GetByLiked(ctx context.Context, userID int64) ([]model.Post, error) {
	posts := make([]model.Post, 0)

	tx, err := p.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return nil, err
	}

	query := `
	SELECT Posts.id, Users.name, Posts.create_att, Posts.title, Posts.content,
    	(select count(*) FROM Reactions_Posts where post_id = Posts.id and type = 1) AS likes,
    	(select count(*) FROM Reactions_Posts where post_id = Posts.id and type =  -1)  AS dislikes,
    	(SELECT COUNT(*) FROM Comments WHERE Comments.post_id = Posts.id) AS sum_comments
    FROM Posts
	JOIN Users ON Posts.user_id = Users.id
    JOIN Reactions_Posts on Reactions_Posts.post_id = Posts.id
    WHERE Reactions_Posts.type = 1 AND Reactions_Posts.user_id = $1
	ORDER by Posts.create_att DESC;`

	rows, err := tx.QueryContext(ctx, query, userID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	queryCategory := `
	SELECT  Categories.name
	FROM Categories_Posts 
	JOIN Categories
		ON Categories_Posts.category_id = Categories.id
	WHERE post_id = ?;`

	for rows.Next() {
		p := model.Post{}
		var timeStr string
		err := rows.Scan(&p.ID, &p.UserName, &timeStr, &p.Title, &p.Content, &p.Like, &p.DisLike, &p.SumComments)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		p.CreateAtt, err = time.Parse("2006-01-02 15:04:05.999999999-07:00", timeStr)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		rows, err := tx.QueryContext(ctx, queryCategory, p.ID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		for rows.Next() {
			var category string

			err := rows.Scan(&category)
			if err != nil {
				tx.Rollback()
				return nil, err
			}

			p.Categories = append(p.Categories, category)
		}

		posts = append(posts, p)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return posts, nil
}
