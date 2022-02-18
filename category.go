package finance

import (

)

type Category struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" binding:"required"`
	Price int `json:"price" binding:"required"`
}

type Category_payments struct {
	Id     int
	CategoryId int
	PaymentId int
}
type Payments struct {
	Id   int    `json:"id" db:"id"`
	Type string `json:"type" db:"type" binding:"required"`
	Date string	`json:"date" db:"date"`
}
type UpdateCategoryInput struct {
	Name     *string `json:"name"`
	Price	 *int `json:"price"`
}
type GetAllPayments struct {
	Id   int    `json:"id"`
	Type string `json:"type" `
	Date string	`json:"date" `
	Name  string `json:"name" `
	Price int `json:"price" `
}
type UpdatePaymentInput struct {
	Type     string `json:"type"`
}