package dot

import (
	"net/url"
)

type UserSignUp struct {
	Email           string
	Name            string
	Password        string
	ConfirmPassword string
}

func Filling(data url.Values) *UserSignUp {
	return &UserSignUp{
		Email:           data.Get("email"),
		Name:            data.Get("name"),
		Password:        data.Get("password"),
		ConfirmPassword: data.Get("confirm_password"),
	}
}

func (u *UserSignUp) ChecPassword() bool {
	if u.Password == u.ConfirmPassword {
		return true
	} else {
		return false
	}
}
