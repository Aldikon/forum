package post_handlers

import (
	"net/http"

	"project/internal/handlers"
)

func (p *postHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handlers.ErrorPage(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
}
