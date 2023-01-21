package middleware_handlers

import (
	"net/http"

	"project/internal/handlers"
)

func (m *middlewareHandler) PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// log.Println(string(debug.Stack()))
				handlers.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			}
		}()
		next.ServeHTTP(w, req)
	})
}
