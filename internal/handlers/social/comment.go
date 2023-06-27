package handlers

import (
	"errors"
	"log"
	"net/http"

	"forum/internal/model"
)

func (s *social) CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	comment := model.CreateComment{}
	comment.ParseForm(r.Form)
	comment.UserID = r.Context().Value(model.UserID).(int64)

	err = s.social.AddComment(r.Context(), comment)
	switch {
	case errors.As(err, &model.FillingErr):
		log.Println(err)
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
		return
	case err != nil:
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	default:
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	}
}
