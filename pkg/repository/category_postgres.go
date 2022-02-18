package repository

import (
	"fmt"
	"strings"

	"github.com/zhansul19/qltIntern"

	"github.com/jmoiron/sqlx"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (cp *CategoryPostgres) Create(category finance.Category) (int, error) {
	tx, err := cp.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createCatQuery := fmt.Sprintf("INSERT INTO %s (name,price) VALUES ($1, $2) RETURNING id", categories)
	row := tx.QueryRow(createCatQuery, category.Name, category.Price)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
func (cp *CategoryPostgres) GetCategory() ([]finance.Category, error) {
	var allCategories []finance.Category
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id", categories)
	err := cp.db.Select(&allCategories, query)
	return allCategories, err
}
func (cp *CategoryPostgres) GetCategoryById(id int) (finance.Category, error) {
	var category finance.Category
	query := fmt.Sprintf("SELECT * FROM %s where id=$1", categories)
	err := cp.db.Get(&category, query, id)
	return category, err
}
func (cp *CategoryPostgres) DeleteCategoryById(id int) error {
	query := fmt.Sprintf("DELETE FROM %s where id=$1", categories)
	_, err := cp.db.Exec(query, id)
	return err
}
func (cp *CategoryPostgres) UpdateCategory(id int,category finance.UpdateCategoryInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if category.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *category.Name)
		argId++
	}

	if category.Price != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *category.Price)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query:= fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d",categories,setQuery,argId)
	args=append(args, id)
	_,err:=cp.db.Exec(query,args...)
	return err
}
