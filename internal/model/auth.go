package model

import (
	"net/url"
	"time"
)

type Registration struct {
	Name            string
	Email           string
	Password        string
	ConfirmPassword string
}

func (r *Registration) ParsForm(form url.Values) {
	r.Name = form.Get("name")
	r.Email = form.Get("email")
	r.Password = form.Get("password")
	r.ConfirmPassword = form.Get("confirm_password")
}

func (r *Registration) Validate() error {
	if err := validateName(r.Name); err != nil {
		return err
	}

	if err := validateEmail(r.Email); err != nil {
		return err
	}

	if err := validatePassword(r.Password); err != nil {
		return err
	}

	if r.Password != r.ConfirmPassword {
		return &fillingError{"The passwords don't match "}
	}

	return nil
}

type LogIn struct {
	Email    string
	Password string
}

func (l *LogIn) ParseForm(form url.Values) {
	l.Email = form.Get("email")
	l.Password = form.Get("password")
}

func (l *LogIn) Validate() error {
	if err := validateEmail(l.Email); err != nil {
		return err
	}
	if err := validatePassword(l.Password); err != nil {
		return err
	}

	return nil
}

type Session struct {
	UserID int64
	Token  string
	EndAtt time.Time
}

func (s *Session) Validate() error {
	if s.UserID < 0 {
		return ErrNotFoundUser
	}

	if s.EndAtt.Before(time.Now()) {
		return ErrTokenExpired
	}

	return nil
}

type User struct {
	ID   int64
	Name string
}
