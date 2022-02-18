package service

import (
	"github.com/zhansul19/qltIntern/pkg/repository"
	"github.com/zhansul19/qltIntern"
)

type CategoryService struct {
	repository repository.Category
}
func NewCategoryService(repo repository.Category)*CategoryService{
	return &CategoryService{repository: repo}
}
func (cs* CategoryService)Create(category finance.Category)(int,error){
	return cs.repository.Create(category)
}
func (cs*CategoryService)GetCategory()([]finance.Category,error){
	return cs.repository.GetCategory()
}
func (cs*CategoryService)GetCategoryById(id int)(finance.Category,error){
	return cs.repository.GetCategoryById(id)
}
func (cs*CategoryService)DeleteCategoryById(id int)(error){
	return cs.repository.DeleteCategoryById(id)
}
func (cs*CategoryService)UpdateCategory(id int,category finance.UpdateCategoryInput)(error){
	return cs.repository.UpdateCategory(id,category)
}