package repository

import (
	"database/sql"
	"fmt"

	"project/internal/dot"
	"project/internal/util"
	"project/model"
)

type CategoryRepository interface{}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *categoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (c *categoryRepository) CreateCategoryForPost(cat *dot.CreateCategory) error {
	query := `INSERT INTO Categories_Post (post_id ,category_id ) VALUES (?, ?);`
	temp := make([]string, len(cat.Categories))

	for _, category := range cat.Categories {
		tempcategory, err := c.ReadByNameCategory(category)
		if err != nil {
			err := c.createCategory(category)
			if err != nil {
				return err
			}

			tempcategory, err = c.ReadByNameCategory(category)
			if err != nil {
				return err
			}
		}
		temp = append(temp, tempcategory.Id)
	}
	stmt, err := c.db.Prepare(query)
	if err != nil {
		return err
	}
	for _, id := range temp {
		err := util.StmtExec(stmt, cat.PostId, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *categoryRepository) createCategory(name string) error {
	query := `INSERT INTO Categories (name) VALUES ('?');`
	stmt, err := c.db.Prepare(query)
	if err != nil {
		return err
	}
	return util.StmtExec(stmt, name)
}

func (c *categoryRepository) ReadByIdCategory(id string) (*model.Category, error) {
	query := `SELECT * FROM Categories WHERE id = ?;`
	stmt, err := c.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	var category *model.Category
	err = stmt.QueryRow(id).Scan(category.Id, category.Name)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *categoryRepository) ReadByNameCategory(name string) (*model.Category, error) {
	query := `SELECT * FROM Categories WHERE name = ?;`
	stmt, err := c.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	var category *model.Category
	err = stmt.QueryRow(name).Scan(category.Id, category.Name)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *categoryRepository) ReadByPostIdCategory(postId string) ([]*model.Category, error) {
	query := `SELECT * FROM Categories_Post WHERE post_id = ?;`
	stmt, err := c.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	var categories []*model.Category

	rows, err := stmt.Query(postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category *model.Category
		err := rows.Scan(category.Id, category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	if len(categories) < 1 {
		return nil, fmt.Errorf("%s post_id not have category", postId)
	}
	return categories, nil
}

func (c *categoryRepository) DeleteCategory(id string) error {
	query := `DELETE FROM Categories WHERE id= ?;`
	stmt, err := c.db.Prepare(query)
	if err != nil {
		return err
	}
	return util.StmtExec(stmt, id)
}
