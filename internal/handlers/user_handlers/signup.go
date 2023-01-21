package user_handlers

import (
	"net/http"

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

	err := handlers.TemplatePage.ExecuteTemplate(w, "signup.html", nil)
	if err != nil {
		handlers.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}
