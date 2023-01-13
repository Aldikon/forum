package controllers

import (
	"html/template"
	"net/http"

	"project/internal/model"
	"project/internal/service"
	"project/internal/util"
)

type SingInController struct {
	ser service.UserService
}

func NewSingInController(s service.UserService) *SingInController {
	return &SingInController{
		ser: s,
	}
}

func (s *SingInController) SignIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		signInGet(w, r, s.ser)
	case http.MethodPost:
		signInPost(w, r, s.ser)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func signInGet(w http.ResponseWriter, r *http.Request, ser service.UserService) {
	temp, err := template.ParseFiles("./ui/template/signin.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)
}

func signInPost(w http.ResponseWriter, r *http.Request, ser service.UserService) {
	// r.ParseForm()
	var user model.User
	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")
	if err := ser.SignInService(user); err != nil {
		bodyOnError := struct {
			Status bool
			Err    string
		}{
			Status: true,
		}
		bodyOnError.Err = "Hello"

		temp, err := template.ParseFiles("./ui/template/signin.html")
		if err != nil {
			// Добавить функцию ошибки
			http.Error(w, "Invalid template", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, err)
		return
	}
	util.DelAllForm(r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
