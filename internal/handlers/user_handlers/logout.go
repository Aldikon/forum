package user_handlers

import (
	"net/http"

	"project/internal/handlers"
)

func (h *userHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handlers.ErrorPage(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	}
	w.Write([]byte("Выход из аккаунта."))
}
