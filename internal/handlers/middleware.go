package handlers

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"runtime/debug"

	"forum/internal/model"
)

type middleware struct {
	service model.AuthService
}

func NewMiddleware(auth model.AuthService) *middleware {
	return &middleware{
		service: auth,
	}
}

func (m *middleware) AuthMiddleware(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Token")
		if err == http.ErrNoCookie {
			// http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			http.Redirect(w, r, "/log-in", http.StatusFound)
			return
		} else if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		id, err := m.service.GetID(r.Context(), cookie.Value)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				http.Redirect(w, r, "/log-in", http.StatusFound)
				return
			}

			if errors.Is(err, model.ErrTokenExpired) {
				http.Redirect(w, r, "/log-in", http.StatusFound)
				return
			}

			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), model.UserID, id))

		http.HandlerFunc(next).ServeHTTP(w, r)
	})
}

func (m *middleware) OptionalAuthMiddleware(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Token")
		if err != nil && err != http.ErrNoCookie {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		var id int64

		if err == http.ErrNoCookie {
			id = 0
		} else {
			id, err = m.service.GetID(r.Context(), cookie.Value)
			if err != nil && !errors.Is(err, sql.ErrNoRows) && !errors.Is(err, model.ErrTokenExpired) {
				log.Println(err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

		}

		r = r.WithContext(context.WithValue(r.Context(), model.UserID, id))

		http.HandlerFunc(next).ServeHTTP(w, r)
	})
}

func (m *middleware) NotAuthMimiddleware(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("Token")
		if err != nil && err != http.ErrNoCookie {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		http.HandlerFunc(next).ServeHTTP(w, r)
	})
}

func (m *middleware) RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				log.Println(string(debug.Stack()))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
