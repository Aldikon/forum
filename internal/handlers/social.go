package handlers

import (
	"net/http"

	handlers "forum/internal/handlers/social"
	"forum/internal/model"
)

func InitSocial(mux *http.ServeMux, m *middleware, s model.SocialService, a model.AuthService) {
	handlers := handlers.NewSocial(s, a)

	mux.Handle("/", m.OptionalAuthMiddleware(handlers.Index))
	mux.Handle("/post/", m.OptionalAuthMiddleware(handlers.GetPost))
	mux.Handle("/post", m.AuthMiddleware(handlers.CreatePost))
	mux.Handle("/comment", m.AuthMiddleware(handlers.CreateComment))
	mux.Handle("/reac-post", m.AuthMiddleware(handlers.ReactionPost))
	mux.Handle("/reac-comment", m.AuthMiddleware(handlers.ReactionComment))
}
