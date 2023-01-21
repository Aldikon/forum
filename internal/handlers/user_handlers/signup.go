package user_handlers

import (
	"net/http"

	"project/internal/dot"
	"project/internal/handlers"
)

func (h *userHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.signUpGet(w, r)
	case http.MethodPost:
		h.signUpPost(w, r)
	default:
		handlers.ErrorPage(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	}
}

func (h *userHandler) signUpGet(w http.ResponseWriter, r *http.Request) {
	err := handlers.TemplatePage.ExecuteTemplate(w, "signup.html", nil)
	if err != nil {
		handlers.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

func (h *userHandler) signUpPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	err := h.userService.SignUp(dot.FillingUserSignUp(r.PostForm))
	if err == nil {
		cleenRequest(r)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	err = handlers.TemplatePage.ExecuteTemplate(w, "signup.html", err.Error())
	if err != nil {
		handlers.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

func cleenRequest(r *http.Request) {
	for key := range r.Form {
		delete(r.Form, key)
	}
	for key := range r.PostForm {
		delete(r.PostForm, key)
	}
}
