package service

import (
	"bookstore-manager/model"
	"bookstore-manager/repository"
)

type CategoryService struct {
	CategoryDAO *repository.CategoryDAO
}

// GetCategoryList 获取所有分类
func (c *CategoryService) GetCategoryList() ([]*model.Category, error) {
	return c.CategoryDAO.GetCategoryList()
}

func NewCategoryService() *CategoryService {
	return &CategoryService{CategoryDAO: repository.NewCategoryDAO()}
}
