package handlers

import (
	"log"
	"net/http"

	"github.com/Aldikon/forum/internal/model"
	"github.com/Aldikon/forum/internal/template"
)

func (s *social) Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	userID := model.GetID(r.Context())

	filter := r.URL.Query().Get("filter")

	userName, err := s.auth.GetByID(r.Context(), userID)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	categories, err := s.social.GetCategoryAll(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	posts := make([]model.Post, 0)

	if filter != "" {
		posts, err = s.social.GetPostAllFilter(r.Context(), userID, filter)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	} else {
		posts, err = s.social.GetPostAll(r.Context())
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	index := model.IndexPage{
		User: model.User{
			ID:   userID,
			Name: userName,
		},
		Categories: categories,
		Posts:      posts,
	}

	err = template.Page["index"].Execute(w, index)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
