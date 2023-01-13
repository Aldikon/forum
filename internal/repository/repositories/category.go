package repositories

import "database/sql"

type categoriRepository struct {
	db *sql.DB
}

func NewCategoriRepository(db *sql.DB) *categoriRepository {
	return &categoriRepository{
		db: db,
	}
}

func (r *categoriRepository) CreateCategory(name string) error {
	records := `
	INSERT INTO Categories (name)
	VALUES (?)`

	stmt, err := r.db.Prepare(records)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return nil
}
