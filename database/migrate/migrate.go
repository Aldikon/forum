package migrate

// СОЗДАНИЕ ТАБЛИЦ ДЛЯ БАЗЫ ДАННЫХ

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTable(db *sql.DB) error {
	path := "../../database/migrate/table_sql"

	// ПРОЧИТКА ВСЕЙ ДИРЕКТОРИИ
	dir, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range dir {
		info, err := file.Info()
		if err != nil {
			return err
		}
		data, err := os.ReadFile(fmt.Sprintf("%s/%s", path, info.Name()))
		if err != nil {
			return err
		}
		if _, err := db.Exec(string(data)); err != nil {
			return err
		}
	}
	return nil
}
