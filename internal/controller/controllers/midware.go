// Занимается проверкой Handler перед оснавным блоком.
package controllers

import (
	"net/http"
)

type middlewareController struct{}

func NewMiddlewareController() *middlewareController {
	return &middlewareController{}
}

func (m *middlewareController) SignUpMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
