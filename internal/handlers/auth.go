package handlers

import (
	"net/http"

	"github.com/Aldikon/forum/internal/model"

	handlers "github.com/Aldikon/forum/internal/handlers/auth"
)

func InitAuth(mux *http.ServeMux, m *middleware, a model.AuthService) {
	handlers := handlers.NewAuthorization(a)

	mux.Handle("/sign-up", m.NotAuthMimiddleware(handlers.SignUp))
	mux.Handle("/log-in", m.NotAuthMimiddleware(handlers.LogIn))
	mux.Handle("/log-out", m.AuthMiddleware(handlers.LogOut))
}
