package util

import (
	"database/sql"
	"net/http"

	"project/model"
)

func CleenRequest(r *http.Request) {
	for key := range r.Form {
		delete(r.Form, key)
	}
	for key := range r.PostForm {
		delete(r.PostForm, key)
	}
}

func StmtExec(stmt *sql.Stmt, arg ...any) error {
	res, err := stmt.Exec(arg...)
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return nil
	}
	if row < 1 {
		return model.ErrToCreate
	}
	return nil
}
