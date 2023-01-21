package handlers

import "net/http"

func ErrorPage(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}
