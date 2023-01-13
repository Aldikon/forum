// ЕЩЕ ЕСТЬ ВЕШИ КОТОРЫЙ НУЖНО ПОСМОТРЕТЬ НО УЖЕ МОЖНО ИСПОЛЬЗОВАТЬ
package repositories

import (
	"database/sql"
	"errors"

	"project/internal/model"
)

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *commentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) CreateComment(comment model.Comment) error {
	var records string
	var res sql.Result

	switch comment.ParentId {
	case 0:
		records = `
		INSERT INTO Comments (user_id, post_id, description)
		VALUES (?, ?, ?)`
		stmt, err := r.db.Prepare(records)
		if err != nil {
			return err
		}
		res, err = stmt.Exec(comment.UserId, comment.PostId, comment.Descripton)
		if err != nil {
			return err
		}
	default:
		records = `
		INSERT INTO Comments (user_id, post_id, parent_id, description)
		VALUES (?, ?, ?, ?)`
		stmt, err := r.db.Prepare(records)
		if err != nil {
			return err
		}
		res, err = stmt.Exec(comment.UserId, comment.PostId, comment.ParentId, comment.Descripton)
		if err != nil {
			return err
		}
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		return errors.New("Not created comment")
	}
	return nil
}

func (r *commentRepository) ReadByPostIDComment(post_id int) ([]model.Comment, error) {
	var comments []model.Comment
	var parent sql.NullInt64
	records := `
    SELECT id, user_id, post_id, parent_id, description FROM Comments
    WHERE post_id = ?`
	stmt, err := r.db.Prepare(records)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(post_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment model.Comment
		err = rows.Scan(
			&comment.Id,
			&comment.UserId,
			&comment.PostId,
			&parent,
			&comment.Descripton)
		if err != nil {
			return nil, err
		}
		if parent.Valid {
			comment.ParentId = int(parent.Int64)
		}
		comments = append(comments, comment)
	}
	if len(comments) == 0 {
		return nil, errors.New("Not found comments")
	}
	return comments, nil
}

func (r *commentRepository) ReadByIDComment(post_id int) (model.Comment, error) {
	var comment model.Comment
	var parent sql.NullInt64
	records := `
    SELECT id, user_id, post_id, parent_id, description FROM Comments
    WHERE post_id =?`
	stmt, err := r.db.Prepare(records)
	if err != nil {
		return comment, err
	}
	err = stmt.QueryRow(post_id).Scan(
		&comment.Id,
		&comment.UserId,
		&comment.PostId,
		&parent,
		&comment.Descripton)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return comment, errors.New("Not found comment")
		}
		return comment, err
	}
	if parent.Valid {
		comment.ParentId = int(parent.Int64)
	}
	return comment, nil
}

func (r *commentRepository) UpdateComment(newComment model.Comment) error {
	records := `
	UPDATE Comments
    SET description = ?
    WHERE id = ?`

	stmt, err := r.db.Prepare(records)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(newComment.Descripton, newComment.Id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count <= 0 {
		return errors.New("Not found comment")
	} else if count > 1 {
		return errors.New("Found more than one comment")
	}
	return nil
}

func (r *commentRepository) DeleteComment(comment_id int) error {
	records := `
			DELETE FROM Comments WHERE id = ?`

	stmt, err := r.db.Prepare(records)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(comment_id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count <= 0 {
		return errors.New("Not found post")
	}
	return err
}
