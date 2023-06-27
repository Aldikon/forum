package migrate

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
)

var querys []string = []string{
	`CREATE TABLE IF NOT EXISTS Users (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE CHECK(name <> ''),
		email VARCHAR(320) NOT NULL CHECK(length(email) <= 320 AND email LIKE '%_@__%.__%'),
		password TEXT NOT NULL CHECK(password <> '')
	);`,

	`CREATE TABLE IF NOT EXISTS Session (
		id INTEGER PRIMARY KEY,
		user_id INTEGER UNIQUE ,
		token TEXT NOT NULL CHECK(token <> '' AND length(token) = 36 ),
		session_end_time TEXT NOT NULL CHECK(session_end_time <> ''),
		FOREIGN KEY (user_id) REFERENCES Users(id)
		
	);`,

	`CREATE TABLE IF NOT EXISTS Posts (
		id INTEGER PRIMARY KEY,
		user_id INTEGER NOT NULL CHECK(user_id > 0),
		create_att TEXT  NOT NULL CHECK(create_att <> ''),
		title TEXT NOT NULL CHECK(length(title) >= 1 AND title <> ''),
		content TEXT NOT NULL CHECK(length(content) >= 1 AND content <> ''),
		FOREIGN KEY (user_id) REFERENCES Users(id)
	);`,

	`CREATE TABLE IF NOT EXISTS Comments (
		id INTEGER PRIMARY KEY,
		user_id INTEGER NOT NULL CHECK(user_id > 0),
		post_id INTEGER NOT NULL CHECK(post_id > 0),
		create_att TEXT  NOT NULL CHECK(create_att <> ''),
		parent_id INTEGER REFERENCES Comments(id) ON DELETE SET NULL,
		content TEXT NOT NULL CHECK(length(content) >= 1 AND content <> ''),
		FOREIGN KEY (user_id) REFERENCES Users(id), 
		FOREIGN KEY (post_id) REFERENCES Posts(id)
	);`,

	`CREATE TABLE IF NOT EXISTS Categories (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE CHECK(name <> '')
	);`,

	`CREATE TABLE IF NOT EXISTS Categories_Posts (
		post_id INTEGER NOT NULL CHECK(post_id > 0),
		category_id INTEGER NOT NULL CHECK(category_id > 0),
		FOREIGN KEY (post_id) REFERENCES Posts(id),
		FOREIGN KEY (category_id) REFERENCES Categories(id)
	);`,

	`CREATE TABLE IF NOT EXISTS Reactions_Posts (
		id INTEGER PRIMARY KEY,
		user_id INTEGER NOT NULL CHECK(user_id > 0),
		post_id INTEGER NOT NULL CHECK(post_id > 0),
		type INTEGER NOT NULL CHECK(type = -1 OR type = 1 OR type = 0), 
		FOREIGN KEY (user_id) REFERENCES Users(id),
		FOREIGN KEY (post_id) REFERENCES Posts(id),
		CONSTRAINT unique_user_post_reaction UNIQUE(user_id, post_id)
	);`,

	`CREATE TABLE IF NOT EXISTS Reactions_Comments (
		id INTEGER PRIMARY KEY,
		user_id INTEGER NOT NULL CHECK(user_id > 0),
		comment_id INTEGER NOT NULL CHECK(comment_id > 0),
		type INTEGER NOT NULL CHECK(type = -1 OR type = 1 OR type = 0),
		FOREIGN KEY (user_id) REFERENCES Users(id),
		FOREIGN KEY (comment_id) REFERENCES Comments(id),
		CONSTRAINT unique_user_comment_reaction UNIQUE(user_id, comment_id)
	);`,
}

var datas []string = []string{
	`INSERT INTO Categories (name) VALUES (
		'Alem'
	);`,
	`INSERT INTO Categories (name) VALUES (
		'Go'
	);`,
	`INSERT INTO Categories (name) VALUES (
		'Rust'
	);`,
	`INSERT INTO Categories (name) VALUES (
		'JS'
	);`,
}

func Migrate(ctx context.Context, db *sql.DB) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for _, query := range querys {
		_, err := tx.ExecContext(ctx, query)
		if err != nil {
			return err
		}
	}
	for _, query := range datas {
		_, err := tx.ExecContext(ctx, query)
		if err != nil {
			var sqliteErr sqlite3.Error
			if errors.As(err, &sqliteErr) {
				if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
					continue
				}
			}
			return err
		}
	}

	return tx.Commit()
}
