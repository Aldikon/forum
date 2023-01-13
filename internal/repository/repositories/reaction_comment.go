// МЫСЫЛЬ ДОБАВИТЬ ЗАПРОС НА СПИСОК ВСЕХ ЗАЛАЙКАНЫХ КОМЕНТОВ НА ОДНОМ ПОСТЕ
package repositories

import (
	"database/sql"
	"errors"
)

type reactionCommentRepository struct {
	db *sql.DB
}

func NewReactionCommentRepository(db *sql.DB) *reactionCommentRepository {
	return &reactionCommentRepository{
		db: db,
	}
}

func (r *reactionCommentRepository) CreateTheCommentReaction(user_id, comment_id, typ int) error {
	records := `
	INSERT INTO Comment_Likes (user_id, comment_id, type)
	VALUES (?,?,?)`
	stmt, err := r.db.Prepare(records)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(user_id, comment_id, typ)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count <= 0 {
		return errors.New("Not created Reaction on Comment")
	}
	return nil
}

func (r *reactionCommentRepository) ReadAllCommentReaction(comment_id int) ([]int, error) {
	records := `
	SELECT type FROM Comment_Likes
	WHERE comment_id =?`

	stmt, err := r.db.Prepare(records)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(comment_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var types []int
	for rows.Next() {
		var typ int
		err = rows.Scan(&typ)
		if err != nil {
			return nil, err
		}
		types = append(types, typ)
	}
	if len(types) == 0 {
		return nil, errors.New("Not found Reaction on Comment")
	}
	return types, nil
}

func (r *reactionCommentRepository) UpdateTheCommentReaction(comment_id, typ int) error {
	if typ < -1 || typ > 1 {
		return errors.New("Invalid type")
	}
	records := `
    UPDATE Comment_Likes
    SET type =?
    WHERE comment_id =?`
	stmt, err := r.db.Prepare(records)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(typ, comment_id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Not updated Reaction on Comment")
	} else if count > 1 {
		return errors.New("More updated Reaction on Comment")
	}
	return nil
}

func (r *reactionCommentRepository) DeleteTheCommentReaction(comment_id int) error {
	records := `
		DELETE FROM Comment_Likes
		WHERE comment_id =?`
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
	if count <= 0 {
		return errors.New("Comment not found!")
	}
	return err
}
