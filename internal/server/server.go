package server

import (
	"net/http"

	"project/config"
)

func NewServer(mux http.Handler) *http.Server {
	return &http.Server{
		Addr:    config.C.Server.Port,
		Handler: mux,
	}
}
