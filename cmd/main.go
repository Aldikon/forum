package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/Aldikon/forum/internal/handlers"
	"github.com/Aldikon/forum/internal/repository"
	"github.com/Aldikon/forum/internal/service"
	"github.com/Aldikon/forum/internal/template"
	"github.com/Aldikon/forum/migrate"
	"github.com/Aldikon/forum/pkg/sqlite"
)

const adr = ":8080"

func dbInit(path string) (*sql.DB, error) {
	db, err := sqlite.Connect(path)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	err = migrate.Migrate(ctx, db)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	err := template.InitTemplate("./static/html")
	if err != nil {
		log.Fatalln(err)
	}

	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)

	db, err := dbInit("./db/data.sqlite")
	if err != nil {
		log.Fatalf("Connect DB: %v", err)
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	authR := repository.NewAuthorization(db)
	postR := repository.NewPost(db)
	commmentR := repository.NewComment(db)
	reactionR := repository.NewReactiona(db)

	authS := service.NewAuthorization(authR)
	socialS := service.NewSocial(postR, commmentR, reactionR)

	midd := handlers.NewMiddleware(authS)

	handlers.InitAuth(mux, midd, authS)
	handlers.InitSocial(mux, midd, socialS, authS)

	server := http.Server{
		Addr:         adr,
		Handler:      midd.RecoverPanic(mux),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	log.Printf("listen server on http://localhost%s port", adr)
	err = server.ListenAndServe()
	log.Fatalln(err)
}
