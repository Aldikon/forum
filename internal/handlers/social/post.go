package handlers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/Aldikon/forum/internal/model"
	"github.com/Aldikon/forum/internal/template"
)

func (s *social) CreatePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.createPostGet(w, r)
	case http.MethodPost:
		s.createPostPost(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (s *social) createPostGet(w http.ResponseWriter, r *http.Request) {
	err := template.Page["create_post"].Execute(w, nil)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (s *social) createPostPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	post := model.CreatePost{}
	post.ParseForm(r.Form)

	post.UserID = model.GetID(r.Context())

	err = s.social.AddPost(r.Context(), post)
	switch {
	case errors.As(err, &model.FillingErr):
		log.Println(err)
		err := template.Page["create_post"].Execute(w, err)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	case err != nil:
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	default:
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (s *social) GetPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	postID, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	userID := model.GetID(r.Context())

	userName, err := s.auth.GetByID(r.Context(), userID)
	if err != nil {
		return
	}

	post, err := s.social.GetPostByID(r.Context(), int64(postID))
	if err != nil {
		log.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			http.NotFound(w, r)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	post.Comments, err = s.social.GetCommentByPostID(r.Context(), int64(postID))
	if err != nil {
		// ...
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	page := model.PostPage{
		User: model.User{
			ID:   userID,
			Name: userName,
		},
		Post: post,
	}

	err = template.Page["post"].Execute(w, page)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
