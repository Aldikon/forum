package user_handlers

import (
	"net/http"
	"time"

	"project/internal/dot"
	"project/internal/handlers"

	"github.com/gofrs/uuid"
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
	r.ParseForm()
	err := h.userService.LogIn(dot.FillingUserLogIn(r.PostForm))
	if err == nil {

		c := http.Cookie{}
		c.Name = "sesseans_token"
		c.MaxAge = int(time.Minute * 5)
		u, err := uuid.NewV4()
		if err != nil {
			handlers.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		c.Value = u.String()
		http.SetCookie(w, &c)
	}

	err = handlers.TemplatePage.ExecuteTemplate(w, "login.html", err)
	if err != nil {
		handlers.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}
