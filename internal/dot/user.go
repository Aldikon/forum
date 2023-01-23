package dot

import (
	"net/url"
	"strings"
)

type UserSignUp struct {
	Email           string
	Name            string
	Password        string
	ConfirmPassword string
}

func FillingUserSignUp(data url.Values) *UserSignUp {
	return &UserSignUp{
		Email:           strings.ToLower(data.Get("email")),
		Name:            data.Get("name"),
		Password:        data.Get("password"),
		ConfirmPassword: data.Get("confirm_password"),
	}
}

type UserLogIn struct {
	Email    string
	Password string
}

func FillingUserLogIn(data url.Values) *UserLogIn {
	return &UserLogIn{
		Email:    strings.ToLower(data.Get("email")),
		Password: data.Get("password"),
	}
}
