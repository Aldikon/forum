package user_handlers

import (
	"net/http"

	"project/internal/handlers"
)

func (h *userHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.logInGet(w, r)
	case http.MethodPost:
		h.logInPost(w, r)
	default:
		handlers.ErrorPage(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	}
}

func (h *userHandler) logInGet(w http.ResponseWriter, r *http.Request) {
	err := handlers.TemplatePage.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		handlers.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

func (h *userHandler) logInPost(w http.ResponseWriter, r *http.Request) {
	err := handlers.TemplatePage.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		handlers.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}
