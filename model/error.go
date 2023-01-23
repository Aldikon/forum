package model

import "errors"

var ErrToCreate = errors.New("impossible to create")

type errorPage struct {
	Code    int
	Message string
}

func NewErrorPage(code int, m string) *errorPage {
	return &errorPage{
		Code:    code,
		Message: m,
	}
}
