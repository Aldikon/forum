package handlers

import (
	"net/http"

	"project/internal/dot"
)

func ErrorPage(w http.ResponseWriter, code int, message string) {
	if err := TemplatePage.ExecuteTemplate(w, "error.html", dot.NewErrorDot(code, message)); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
