package handlers

import (
	"log"
	"net/http"

	"github.com/Aldikon/forum/internal/model"
)

func (a *auth) LogOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	userID := model.GetID(r.Context())

	err := a.AuthService.LogOut(r.Context(), userID)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:   "Token",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusFound)
}
