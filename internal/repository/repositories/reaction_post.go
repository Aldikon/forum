package repositories

import (
	"database/sql"
	"errors"
)

type reactionPostRepository struct {
	db *sql.DB
}

func NewPostReactionRepository(db *sql.DB) *reactionPostRepository {
	return &reactionPostRepository{
		db: db,
	}
}

func (r *reactionPostRepository) CreateThePostReaction(user_id, post_id, typ int) error {
	records := `
	INSERT INTO Post_Likes (user_id, post_id, type)
	VALUES (?,?,?)`
	stmt, err := r.db.Prepare(records)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(user_id, post_id, typ)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count <= 0 {
		return errors.New("Not created Reaction on post")
	}
	return nil
}

func (r *reactionPostRepository) ReadAllPostReaction(post_id int) ([]int, error) {
	records := `
	SELECT type FROM Post_Likes
	WHERE post_id =?`
	stmt, err := r.db.Prepare(records)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(post_id)
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
		return nil, errors.New("Not found Reaction on post")
	}
	return types, nil
}

func (r *reactionPostRepository) UpdateThePostReaction(post_id, typ int) error {
	if typ < -1 || typ > 1 {
		return errors.New("Invalid type")
	}
	records := `
    UPDATE Post_Likes
    SET type =?
    WHERE post_id =?`
	stmt, err := r.db.Prepare(records)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(typ, post_id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Not updated Reaction on post")
	} else if count > 1 {
		return errors.New("More updated Reaction on post")
	}
	return nil
}

func (r *reactionPostRepository) DeleteThePostReaction(post_id int) error {
	records := `
		DELETE FROM Post_Likes
		WHERE post_id =?`
	stmt, err := r.db.Prepare(records)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(post_id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if count <= 0 {
		return errors.New("post not found!")
	}
	return err
}
