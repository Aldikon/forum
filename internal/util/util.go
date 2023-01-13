package util

import (
	"net/http"
	"regexp"
)

func CheckEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}

func CheckUserName(userName string) bool {
	re := regexp.MustCompile("^[A-Za-z0-9]+([A-Za-z0-9]*|[._-]?[A-Za-z0-9]+)*$")
	return re.MatchString(userName)
}

func ValidPasswords(pas string) bool {
	re := regexp.MustCompile(`^[A-Za-z0-9]+`)
	return re.MatchString(pas)
}

func DelAllForm(r *http.Request) {
	for key := range r.Form {
		delete(r.Form, key)
	}
	for key := range r.PostForm {
		delete(r.PostForm, key)
	}
}
