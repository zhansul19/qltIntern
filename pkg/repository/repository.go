package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zhansul19/qltIntern"

)

type Category interface {
	Create( category finance.Category) (int, error)
	GetCategory()([]finance.Category,error)
	GetCategoryById(id int)(finance.Category,error)
	DeleteCategoryById(id int)(error)
	UpdateCategory(id int,category finance.UpdateCategoryInput)(error)

}
type Payment interface {
	Create(categoryId int,payment finance.Payments) (int, error)
	GetAll(categoryId int)([]finance.GetAllPayments,error)
	GetPaymentById(id int) (finance.Payments, error)
	DeletePaymentById(id int) error
	UpdatePayments(id int, payment finance.UpdatePaymentInput) error

}
type Repository struct {
	Category
	Payment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Category:NewCategoryPostgres(db),
		Payment: NewPaymentPostgres(db),
	}
}