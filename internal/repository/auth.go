package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Aldikon/forum/internal/model"
	"github.com/mattn/go-sqlite3"
)

type auth struct {
	db *sql.DB
}

func NewAuthorization(db *sql.DB) *auth {
	return &auth{
		db: db,
	}
}

func (r *auth) AddUser(ctx context.Context, user model.Registration) error {
	query := `
	INSERT INTO USERS (name, email, password)
	VALUES ($1, $2, $3)`

	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return model.ErrEmail
			}
		}
		return err
	}

	return nil
}

func (a *auth) LogIn(ctx context.Context, user model.LogIn) (int64, error) {
	query := `SELECT Users.id FROM Users WHERE email = $1 AND password = $2`

	row := a.db.QueryRowContext(ctx, query, user.Email, user.Password)
	if err := row.Err(); err != nil {
		return 0, err
	}

	var id int64
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (a *auth) AddSession(ctx context.Context, session model.Session) error {
	query := `INSERT OR REPLACE
	INTO Session (user_id, token, session_end_time)
	VALUES ($1, $2, $3)`

	res, err := a.db.ExecContext(ctx, query, session.UserID, session.Token, session.EndAtt)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count <= 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (a *auth) GetID(ctx context.Context, token string) (model.Session, error) {
	session := model.Session{}

	query := `SELECT user_id, token, session_end_time 
	FROM Session WHERE token = $1;`

	row := a.db.QueryRow(query, token)

	var endAttStr string

	err := row.Scan(&session.UserID, &session.Token, &endAttStr)
	if err != nil {
		return session, err
	}
	session.EndAtt, err = time.Parse("2006-01-02 15:04:05.999999999-07:00", endAttStr)
	if err != nil {
		return session, err
	}

	return session, nil
}

func (a *auth) GetByID(ctx context.Context, userID int64) (string, error) {
	var userName string
	query := `SELECT Users.name FROM Users WHERE Users.id = $1;`

	row := a.db.QueryRowContext(ctx, query, userID)

	err := row.Scan(&userName)
	if err != nil {
		return "", err
	}
	return userName, nil
}

func (a *auth) DeleteSession(ctx context.Context, userID int64) error {
	query := `DELETE FROM Session WHERE user_id = $1;`

	_, err := a.db.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}

	return nil
}
