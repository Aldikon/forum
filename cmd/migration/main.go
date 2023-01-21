package main

// ФАЙЛ ДЛЯ ИЗМИНЕНИЕ И СБРОСА БАЗЫ ДАННЫХ

import (
	"log"
	"os"

	migrate "project/db"
	"project/pkg/sqlite"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) <= 1 || len(os.Args) >= 3 {
		log.Fatal("Usage: go run change_db.go <argument>")
	}
	flag := os.Args[1]
	db, err := sqlite.Connect("./db/forum.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	switch flag {
	case "create":
		if err := migrate.CreateTable(db); err != nil {
			log.Fatalln(err)
		}
		log.Println("Successful")
	case "drop":
		if err := sqlite.DropAllDB(db); err != nil {
			log.Fatalln(err)
		}
		log.Println("Successful")
	default:
		log.Fatalf("%s: unknown flag. Use: 'create' or 'drop'", flag)
	}
}
