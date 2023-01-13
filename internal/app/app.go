package app

import (
	"database/sql"
	"log"
	"net/http"

	"project/internal/controller"
	"project/internal/repository"
	"project/internal/service"
	sqlite "project/package/sqllite"
)

const (
	port = ":8989"
)

func buildMux(con controller.Controller) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/signup", con.SignUpMiddleware(http.HandlerFunc(con.SignUp)))
	mux.HandleFunc("/signin", con.SignIn)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Build mux")
	return mux
}

func layering(db *sql.DB) controller.Controller {
	repo := repository.NewRepository(db)
	log.Println("Get repository")

	service := service.NewService(repo)
	log.Println("Get service")

	controller := controller.NewController(service)
	log.Println("Get controller")

	return controller
}

func Run(server *http.Server) error {
	db, err := sqlite.Connect("./database/forum.sqlite")
	if err != nil {
		return err
	}
	log.Println("Connecting db.")

	controller := layering(db)

	mux := buildMux(controller)

	server.Handler = mux

	log.Printf("Starting listening on http://localhost%s ", server.Addr)

	return server.ListenAndServe()
}
