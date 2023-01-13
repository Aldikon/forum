package model

import (
	"errors"

	"project/internal/util"
)

// Идея добавить поле день рождение. Возможно не нужно. Форум может обойтись без твоего дня рождения
// Идея добавить поле дата регистрации.
type User struct {
	Id       int
	UserName string
	Email    string
	Password string
}

// Идея добавить поле создание поста.
type Post struct {
	Id         int
	Title      string
	Descripton string
	UserId     int
	Category   []string
}

// Идея добавить поле создания коммента.
type Comment struct {
	Id         int
	UserId     int
	PostId     int
	ParentId   int
	Descripton string
}

// USER METHODS ----------------------------------------------------------------
func (u *User) IsEmpty() bool {
	if u.UserName == "" || u.Email == "" || u.Password == "" {
		return true
	}
	return false
}

func (u *User) CheckInput(confirmPassword string) error {
	if u.IsEmpty() {
		return errors.New("Please fill in the input fields!")
	}
	if !(util.CheckEmail(u.Email)) {
		return errors.New("Not a valid email address!")
	}
	if !(util.CheckUserName(u.UserName)) {
		return errors.New("Not a valid user name!")
	}
	if confirmPassword != u.Password {
		return errors.New("Passwords don't match!")
	}
	if !(util.ValidPasswords(u.Password)) {
		return errors.New("Not a valid passwords!")
	}
	return nil
}

// COMMENTS METHODS ----------------------------------------------------------------
func (p *Comment) IsHasParentComment() bool {
	return p.ParentId <= 0
}

func (p *Comment) IsEmpty() bool {
	return p.Descripton == ""
}
