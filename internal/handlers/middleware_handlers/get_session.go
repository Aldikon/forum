package middleware_handlers

import (
	"errors"
	"net/http"

	"project/internal/handlers"
)

func (m *middlewareHandler) MiddlewareGetSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("session")
		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				handlers.ErrorPage(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			default:
				handlers.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			}
			return
		}
		// Отправить запрос на сеанс для проверки
		next.ServeHTTP(w, r)
	})
}
