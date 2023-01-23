package repository

import (
	"database/sql"
	"errors"

	"project/internal/dot"
	"project/model"
)

type UserRepository interface {
	CreateUser(user *dot.UserSignUp) error
	ReadToRegisterUser(*dot.UserLogIn) error
	ReadToIdUser(id string) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUser(user *dot.UserSignUp) error {
	query := `
	INSERT INTO Users (name, email, password) 
	VALUES (?,?,?)`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return model.ErrToCreate
	}
	return nil
}

func (u *userRepository) ReadToRegisterUser(user *dot.UserLogIn) error {
	query := `	SELECT * FROM Users WHERE email = ? AND password = ?`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows := stmt.QueryRow(user.Email, user.Password)
	err = rows.Err()
	if err != nil {
		return errors.New("User not found!")
	}
	return nil
}

func (u *userRepository) ReadToIdUser(id string) (*model.User, error) {
	query := `	SELECT * FROM Users WHERE id = ? `
	stmt, err := u.db.Prepare(query)
	if err != nil {
		return nil, nil
	}
	defer stmt.Close()

	rows := stmt.QueryRow(id)
	var user *model.User

	err = rows.Scan(user.Id, user.Name, user.Email, user.Password)
	if err != nil {
		return nil, errors.New("User not found!")
	}

	return user, nil
}
