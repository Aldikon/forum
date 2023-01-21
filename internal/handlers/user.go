package handlers

import (
	"net/http"

	"project/internal/service"
)

type UserHandler interface {
	Login(http.ResponseWriter, *http.Request)
	Logout(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

// url = /login ------------------------------------------------------------------------------------
func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.loginGet(w, r)
	case http.MethodPost:
		h.loginPost(w, r)
	default:
		ErrorPage(w, http.StatusMethodNotAllowed)
	}
}

func (h *userHandler) loginGet(w http.ResponseWriter, r *http.Request) {
	err := templatePage.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
	}
}

func (h *userHandler) loginPost(w http.ResponseWriter, r *http.Request) {
	err := templatePage.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
	}
}

// url = /logout ------------------------------------------------------------------------------------
func (h *userHandler) Logout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.logoutGet(w, r)
	case http.MethodPost:
		h.logoutPost(w, r)
	default:
		ErrorPage(w, http.StatusMethodNotAllowed)
	}
}

func (h *userHandler) logoutGet(w http.ResponseWriter, r *http.Request) {
	err := templatePage.ExecuteTemplate(w, "logout.html", nil)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
	}
}

func (h *userHandler) logoutPost(w http.ResponseWriter, r *http.Request) {
	err := templatePage.ExecuteTemplate(w, "logout.html", nil)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
	}
}

// url = /Profile ------------------------------------------------------------------------------------
func (h *userHandler) Profile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.profileGet(w, r)
	case http.MethodPost:
		h.profilePost(w, r)
	default:
		ErrorPage(w, http.StatusMethodNotAllowed)
	}
}

func (h *userHandler) profileGet(w http.ResponseWriter, r *http.Request) {
	err := templatePage.ExecuteTemplate(w, "profile.html", nil)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
	}
}

func (h *userHandler) profilePost(w http.ResponseWriter, r *http.Request) {
	err := templatePage.ExecuteTemplate(w, "profile.html", nil)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
	}
}

// url = /registration ------------------------------------------------------------------------------------
func (h *userHandler) Registration(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.registrationGet(w, r)
	case http.MethodPost:
		h.registrationPost(w, r)
	default:
		ErrorPage(w, http.StatusMethodNotAllowed)
	}
}

func (h *userHandler) registrationGet(w http.ResponseWriter, r *http.Request) {
	err := templatePage.ExecuteTemplate(w, "logout.html", nil)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
	}
}

func (h *userHandler) registrationPost(w http.ResponseWriter, r *http.Request) {
	err := templatePage.ExecuteTemplate(w, "logout.html", nil)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
	}
}
