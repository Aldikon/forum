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
	"project/internal/handlers/user_handlers"
	"project/internal/repository"
	"project/internal/server"
	"project/internal/service"
	"project/pkg/sqlite"
)

// todo add config file
type app struct {
	chanErr chan error
}

const (
	urlIndex = "/"

	// url user
	urlLogIn   = "/login"
	urlLogOut  = "/logout"
	urlProfile = "/profile"
	urlSignUp  = "/signup"

	urlPost    = "/post"
	urlLike    = "/like"
	urlCommemt = "/comment"
)

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

func (a *app) build(mux *http.ServeMux, db *sql.DB) {
	// withAccess := http.NewServeMux()
	// withoutAccess := http.NewServeMux()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewService(userRepository)
	userHandler := user_handlers.NewUserHandler(userService)

	// middleware := middleware_handlers.NewMiddlewareHandler()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// user mux
	mux.HandleFunc(urlLogIn, userHandler.LogIn)
	mux.HandleFunc(urlLogOut, userHandler.LogOut)
	mux.HandleFunc(urlProfile, userHandler.Profile)
	mux.HandleFunc(urlSignUp, userHandler.SignUp)

	// mux.Handle("", middleware.PanicRecovery(withAccess))
	// mux.Handle("/", middleware.PanicRecovery(withoutAccess))

	// mux.HandleFunc(urlLike, userHandler.Logout)
	// mux.HandleFunc(urlCommemt, userHandler.Logout)
	// mux.HandleFunc(urlIndex, userHandler.Logout)
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
