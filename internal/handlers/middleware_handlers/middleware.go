package middleware_handlers

import (
	"net/http"
)

type middlewareHandler struct {
	// userService service.UserService
}

func NewMiddlewareHandler() *middlewareHandler {
	return &middlewareHandler{
		// userService: userService,
	}
}

// шаблон
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
