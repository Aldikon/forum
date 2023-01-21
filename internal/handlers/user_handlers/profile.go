package user_handlers

import (
	"net/http"

	"project/internal/handlers"
)

func (h *userHandler) Profile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		handlers.ErrorPage(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	}
	w.Write([]byte("Страничка профиля."))
}
