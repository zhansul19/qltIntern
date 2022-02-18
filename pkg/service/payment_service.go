package service

import (
	finance "github.com/zhansul19/qltIntern"
	"github.com/zhansul19/qltIntern/pkg/repository"
)

type PaymentService struct {
	repository     repository.Payment
	categoryRepo repository.Category
}
func NewPaymentService(repo repository.Payment,categoryRepo repository.Category)*PaymentService{
	return&PaymentService{repository: repo,categoryRepo: categoryRepo}
}
func (ps* PaymentService)Create(categoryId int,payment finance.Payments)(int,error){
	return ps.repository.Create(categoryId,payment)
}
func (ps* PaymentService)GetAll(categoryId int)([]finance.GetAllPayments,error){
	return ps.repository.GetAll(categoryId)
}
func (ps* PaymentService)GetPaymentById(paymentId int)(finance.Payments,error){
	return ps.repository.GetPaymentById(paymentId)
}
func (ps* PaymentService)DeletePaymentById(paymentId int)(error){
	return ps.repository.DeletePaymentById(paymentId)
}
func (ps* PaymentService)UpdatePayments(id int,payment finance.UpdatePaymentInput)(error){
	return ps.repository.UpdatePayments(id,payment)
}