package handlers

import (
	"net/http"

	handlers "forum/internal/handlers/auth"
	"forum/internal/model"
)

func InitAuth(mux *http.ServeMux, m *middleware, a model.AuthService) {
	handlers := handlers.NewAuthorization(a)

	mux.Handle("/sign-up", m.NotAuthMimiddleware(handlers.SignUp))
	mux.Handle("/log-in", m.NotAuthMimiddleware(handlers.LogIn))
	mux.Handle("/log-out", m.AuthMiddleware(handlers.LogOut))
}
