package main

import (
	"log"

	"project/config"
	"project/internal/app"
)

func main() {
	if err := config.ReadConfig(); err != nil {
		log.Fatalln(err)
	}

	myApp := app.NewApp()

	if err := myApp.Run(); err != nil {
		log.Fatalln(err)
	}
}
