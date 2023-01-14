package controllers

import (
	"html/template"
	"net/http"

	"project/internal/model"
	"project/internal/service"
	"project/internal/util"
)

type SingUpController struct {
	ser service.UserService
}

func NewSingUpController(s service.UserService) *SingUpController {
	return &SingUpController{
		ser: s,
	}
}

func (s *SingUpController) SignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		signUpGet(w, r, s.ser)
	case http.MethodPost:
		signUpPost(w, r, s.ser)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func signUpGet(w http.ResponseWriter, r *http.Request, ser service.UserService) {
	temp, err := template.ParseFiles("./ui/template/signup.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)
}

func signUpPost(w http.ResponseWriter, r *http.Request, ser service.UserService) {
	var user model.User
	user.Email = r.FormValue("email")
	user.UserName = r.FormValue("username")
	user.Password = r.FormValue("password")
	if err := ser.SignUpService(user, r.FormValue("confirm_password")); err != nil {
		bodyOnError := struct {
			Status bool
			Err    string
		}{
			Status: true,
		}
		bodyOnError.Err = err.Error()

		temp, err := template.ParseFiles("./ui/template/signup.html")
		if err != nil {
			// Добавить функцию ошибки
			http.Error(w, "Invalid template", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, bodyOnError)
		return
	}
	util.DelAllForm(r)
	http.Redirect(w, r, "/signin", http.StatusSeeOther)
}
