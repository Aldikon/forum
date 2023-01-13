package controllers

import (
	"html/template"
	"log"
	"net/http"

	"project/internal/service"
)

type ErrorController struct {
	ser service.Service
}

func NewErrorController(s service.Service) *ErrorController {
	return &ErrorController{
		ser: s,
	}
}

func (e *ErrorController) Error(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./ui/template/error.html")
	if err != nil {
		log.Println(err)
		return
	}
	temp.Execute(w, nil)
}
