package handlers

import (
	"errors"
	"log"
	"net/http"

	"forum/internal/model"
	"forum/internal/template"
)

func (a *auth) SignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.sigUpGet(w, r)
	case http.MethodPost:
		a.sigUpPost(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (a *auth) sigUpGet(w http.ResponseWriter, r *http.Request) {
	err := template.Page["sign_up"].Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (a *auth) sigUpPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	reg := model.Registration{}
	reg.ParsForm(r.Form)

	err = a.Add(r.Context(), reg)

	switch {
	case errors.As(err, &model.FillingErr):
		log.Println(err)
		err := template.Page["sign_up"].Execute(w, err.Error())
		if err != nil {
			log.Println(err)
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
