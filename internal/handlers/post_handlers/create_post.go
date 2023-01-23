package post_handlers

import (
	"log"
	"net/http"

	"project/internal/dot"
	"project/internal/handlers"
	"project/internal/util"
)

func (p *postHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handlers.ErrorPage(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	err := r.ParseForm()
	if err != nil {
		handlers.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	log.Println(r.Form)
	log.Println(r.PostForm)

	if err := p.postService.CreatePost(dot.FillingCreatePost(r.PostForm)); err != nil {
		handlers.ErrorPage(w, http.StatusInternalServerError, err.Error())
		// handlers.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	util.CleenRequest(r)
	http.Redirect(w, r, "/", http.StatusCreated)
}
