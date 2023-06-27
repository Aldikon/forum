package handlers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Aldikon/forum/internal/model"
	"github.com/Aldikon/forum/internal/template"
)

func (a *auth) LogIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.loginGet(w, r)
	case http.MethodPost:
		a.loginPost(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)

	}
}

func (a *auth) loginGet(w http.ResponseWriter, r *http.Request) {
	err := template.Page["log_in"].Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (a *auth) loginPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	user := model.LogIn{}
	user.ParseForm(r.Form)

	ses, err := a.AuthService.LogIn(r.Context(), user)
	switch {
	case errors.As(err, &model.FillingErr):
		log.Println(err)
		err := template.Page["log_in"].Execute(w, err)
		w.WriteHeader(http.StatusBadRequest)

		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		return
	case err != nil:
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "Token",
		Value:  ses.Token,
		MaxAge: int(time.Hour),
	})

	http.Redirect(w, r, "/", http.StatusFound)
}
