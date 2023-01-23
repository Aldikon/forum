package handlers

import (
	"net/http"

	"project/model"
)

func ErrorPage(w http.ResponseWriter, code int, message string) {
	if err := TemplatePage.ExecuteTemplate(w, "error.html", model.NewErrorPage(code, message)); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
