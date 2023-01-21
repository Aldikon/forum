package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"project/config"
	"project/internal/handlers"
	"project/internal/repository"
	"project/internal/server"
	"project/internal/service"
	"project/pkg/sqlite"
)

// todo add config file
type app struct {
	chanErr chan error
}

func NewApp() *app {
	// to do read config
	// return struct aap
	return &app{
		chanErr: make(chan error),
	}
}

func (a *app) Run() error {
	db, err := sqlite.Connect(config.C.Path.DB)
	if err != nil {
		return err
	}

	if err := handlers.ReadTemplate(config.C.Path.Template); err != nil {
		return err
	}

	mux := http.NewServeMux()
	a.build(mux, db)

	myServer := server.NewServer(mux)

	log.Printf("Starting listener on http://localhost%s", config.C.Server.Port)

	go func() { a.chanErr <- myServer.ListenAndServe() }()

	return a.wait()
}

func (a *app) wait() error {
	// create channel for wait signal
	syscalCh := make(chan os.Signal, 1)
	signal.Notify(syscalCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-syscalCh:
		fmt.Println()
		log.Printf("Stop server...\n")
		return nil
	case err := <-a.chanErr:
		return err
	}
}

const (
	urlIndex   = "/"
	urlLogIn   = "/login"
	urlLogOut  = "/logout"
	urlProfile = "/profile"
	urlPost    = "/post"
	urlLike    = "/like"
	urlCommemt = "/comment"
)

func (a *app) build(mux *http.ServeMux, db *sql.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	fileServer := http.FileServer(http.Dir(config.C.Path.Template))
	mux.Handle("/"+config.C.Path.Static+"/", http.StripPrefix("/"+config.C.Path.Static, fileServer))

	mux.HandleFunc(urlLogIn, userHandler.Login)
	mux.HandleFunc(urlLogOut, userHandler.Logout)
	// mux.HandleFunc(urlProfile, userHandler.Logout)
	// mux.HandleFunc(urlPost, userHandler.Logout)
	// mux.HandleFunc(urlLike, userHandler.Logout)
	// mux.HandleFunc(urlCommemt, userHandler.Logout)
	// mux.HandleFunc(urlIndex, userHandler.Logout)
}
