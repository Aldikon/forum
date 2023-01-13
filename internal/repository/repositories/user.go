package repositories

// ОПИСАНИЕ МЕТОДОВ ОБЪЕКТА ЮЗЕР ДЛЯ РЕПОЗИТОРИЯ

import (
	"database/sql"
	"errors"

	"project/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user model.User) error {
	records := `
	INSERT INTO Users (user_name, email, password) 
	VALUES (?,?,?)`
	stmt, err := r.db.Prepare(records)
	if err != nil {
		return err
	}
	defer stmt.Close()
	count, err := stmt.Exec(user.UserName, user.Email, user.Password)
	if err != nil {
		return err
	}
	rows, err := count.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("User does not exist!")
	}
	return err
}

func (r *UserRepository) ReadToRegisterUser(user model.User) error {
	records := `SELECT * FROM Users WHERE email = ? AND password = ?`
	stmt, err := r.db.Prepare(records)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(user.Email, user.Password)
	if err != nil {
		return err
	}
	defer rows.Close()
	var count uint
	for rows.Next() {
		count++
	}
	if count != 1 {
		return errors.New("User not found!")
	}
	return nil
}

func (r *UserRepository) ReadUser(margin, arg string) (model.User, error) {
	var records string
	var user model.User
	switch margin {
	case "id":
		records = `SELECT * FROM Users WHERE id= ?`
	case "email":
		records = `SELECT * FROM Users WHERE email = ?`
	case "name":
		records = `SELECT * FROM Users WHERE user_name = ?`
	default:
		return user, errors.New("Not correct margin value!")
	}
	if arg == "" {
		return user, errors.New("Empty argument!")
	}

	stmt, err := r.db.Prepare(records)
	if err != nil {
		return user, err
	}
	defer stmt.Close()

	if err := stmt.QueryRow(arg).Scan(&user.Id, &user.UserName, &user.Email, &user.Password); err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(newUser model.User, email string) error {
	records := `
	UPDATE Users
	SET user_name = ?, email = ?, password = ?
	WHERE email =?`
	if newUser.IsEmpty() {
		return errors.New("Empty user!")
	}
	stmt, err := r.db.Prepare(records)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(newUser.UserName, newUser.Email, newUser.Password, email)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count <= 0 {
		return errors.New("User not found!")
	}
	return err
}

func (r *UserRepository) DeleteUser(user model.User) error {
	records := `
	DELETE FROM Users 
	WHERE user_name= ? AND email=? AND password=?`
	if user.IsEmpty() {
		return errors.New("Empty user!")
	}
	stmt, err := r.db.Prepare(records)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.UserName, user.Email, user.Password)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if count <= 0 {
		return errors.New("User not found!")
	}
	return err
}

// post service
// create post

// tag service
// create manyTag repository

// http controller create post
