package repositories

import (
	"database/sql"
)

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}

// func (r *postRepository) CreatePost(post model.Post) error {
// 	var records string
// 	var category sql.NullString
// 	switch post.Category {
// 	case "":
// 		records := `
// 		INSERT INTO Posts (title, description, user_id)
// 		VALUES (?, ?, ?);`
// 		stmt, err := r.db.Prepare(records)
// 		if err != nil {
// 			return err
// 		}
// 		defer stmt.Close()
// 	default:
// 		records := `
// 		INSERT INTO Posts (title, description, category_id, user_id)
// 		VALUES (?,?,?,?);`

// 	}

// 	_, err = stmt.Exec(post.Title, post.Descripton, post.UserId)
// 	return err
// }

// func (r *postRepository) ReadAllPost() ([]model.Post, error) {
// 	records := `
// 	SELECT 	id, title, description, user_id,
// 	(SELECT name FROM Categories WHERE id = category_id) As category_name
// 	FROM Posts;`
// 	stmt, err := r.db.Prepare(records)
// 	if err != nil {
// 		return nil, err
// 	}
// 	rows, err := stmt.Query(records)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var posts []model.Post
// 	for rows.Next() {
// 		var post model.Post
// 		err = rows.Scan(&post.Id, &post.Title, &post.Descripton, &post.UserId)
// 		if err != nil {
// 			return nil, err
// 		}
// 		posts = append(posts, post)
// 	}
// 	if len(posts) == 0 {
// 		return nil, errors.New("Not found posts")
// 	}
// 	return posts, nil
// }

// func (r *postRepository) ReadByUserPost(id int) ([]model.Post, error) {
// 	var posts []model.Post
// 	records := `
// 	SELECT id, title, description, user_id FROM Posts
// 	WHERE user_id = ?;`
// 	stmt, err := r.db.Prepare(records)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()
// 	rows, err := stmt.Query(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var post model.Post
// 		err = rows.Scan(&post.Id, &post.Title, &post.Descripton, &post.UserId)
// 		if err != nil {
// 			return nil, err
// 		}
// 		posts = append(posts, post)
// 	}
// 	if len(posts) == 0 {
// 		return nil, errors.New("Not found posts")
// 	}
// 	return posts, nil
// }

// func (r *postRepository) ReadByIDPost(id int) (model.Post, error) {
// 	var post model.Post
// 	records := `
// 	SELECT id, title, description, user_id
// 	FROM Posts
//     WHERE id =?`
// 	stmt, err := r.db.Prepare(records)
// 	if err != nil {
// 		return post, err
// 	}
// 	defer stmt.Close()
// 	err = stmt.QueryRow(id).Scan(&post.Id, &post.Title, &post.Descripton, &post.UserId)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return post, errors.New("Not found post")
// 		}
// 		return post, err
// 	}
// 	return post, nil
// }

// // НУЖНО ИЗМЕНИТЬ СТРУКТУРУ. В СЛУЧАЕ ИЗМИНЕНИЕ НЕСКОЛЬКИХ ПОСТОВ, НУЖНО
// // ВЕРНУТЬ ПРЕВЕДУШИЕ ЗНАЧЕНИЯ.
// //
// // НО ПО id НЕ МОЖЕТ БЫТЬ БОЛЬШЕ ОБНОГО ОБЪЕКТА.
// func (r *postRepository) UpdatePost(newPost model.Post, post_id int) error {
// 	records := `
// 	UPDATE Posts
//     SET title = ?, description = ?, user_id = ?
// 	WHERE id = ?;`
// 	stmt, err := r.db.Prepare(records)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()
// 	res, err := stmt.Exec(newPost.Title, newPost.Descripton, newPost.UserId, post_id)
// 	if err != nil {
// 		return err
// 	}
// 	count, err := res.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if count <= 0 {
// 		return errors.New("Not found post")
// 	} else if count > 1 {
// 		return errors.New("Found more than one post")
// 	}
// 	return nil
// }

// func (r *postRepository) DeletePost(post_id int) error {
// 	records := `
// 	DELETE FROM Posts WHERE id =?;`
// 	stmt, err := r.db.Prepare(records)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()
// 	res, err := stmt.Exec(post_id)
// 	if err != nil {
// 		return err
// 	}
// 	count, err := res.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if count <= 0 {
// 		return errors.New("Not found post")
// 	}
// 	return err
// }
