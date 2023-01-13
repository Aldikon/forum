package main

import (
	"log"
	"net/http"
	"time"

	"project/internal/app"
)

const port = ":8080"

func main() {
	server := http.Server{
		Addr:         port,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	if err := app.Run(&server); err != nil {
		log.Fatalln(err)
	}
}
