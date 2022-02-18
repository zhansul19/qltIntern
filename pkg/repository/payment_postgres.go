package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/zhansul19/qltIntern"
)

type PaymentPostgres struct {
	db *sqlx.DB
}

func NewPaymentPostgres(db *sqlx.DB)*PaymentPostgres{
	return &PaymentPostgres{db: db}
}
func (pp *PaymentPostgres) Create(categoryId int,payment finance.Payments) (int, error) {
	tx, err := pp.db.Begin()
	if err != nil {
		return 0, err
	}
	var paymentId int
	createPaymentQuery := fmt.Sprintf("INSERT INTO %s (type) values($1) returning id", payments)

	row := tx.QueryRow(createPaymentQuery, payment.Type)
	err = row.Scan(&paymentId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createCategoryPaymentQuery := fmt.Sprintf("INSERT INTO %s(category_id, payment_id)values($1,$2)", category_payments)
	_, err = tx.Exec(createCategoryPaymentQuery, categoryId, paymentId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return paymentId, tx.Commit()
}
func(pp*PaymentPostgres)GetAll(categoryId int)([]finance.GetAllPayments,error){
	var payment []finance.GetAllPayments
	query := fmt.Sprintf("SELECT p.id,p.type,p.date,c.name,c.price FROM %s p INNER JOIN %s cp on cp.payment_id=p.id INNER JOIN %s c on c.id=cp.category_id WHERE c.id = $1 ",
		payments, category_payments, categories)
	if err := pp.db.Select(&payment, query, categoryId); err != nil {
		return nil, err
	}

	return payment, nil
}
func(pp*PaymentPostgres)GetPaymentById(id int) (finance.Payments, error){
	var payment finance.Payments
	query := fmt.Sprintf("SELECT * FROM %s where id=$1", payments)
	err := pp.db.Get(&payment, query, id)
	return payment, err
}
func(pp*PaymentPostgres)DeletePaymentById(id int) error{
	query := fmt.Sprintf("DELETE FROM %s where id=$1", payments)
	_, err := pp.db.Exec(query, id)
	return err
}
func(pp*PaymentPostgres)UpdatePayments(id int, payment finance.UpdatePaymentInput) error{

	

	query:= fmt.Sprintf("UPDATE %s SET type ='%s' WHERE id=$1",payments,payment.Type)
	_,err:=pp.db.Exec(query,id)
	return err
}
