package repository

import (
	"bookstore-manager/global"
	"bookstore-manager/model"

	"gorm.io/gorm"
)

type CategoryDAO struct {
	db *gorm.DB
}

// GetCategoryList 获取所有分类
func (c *CategoryDAO) GetCategoryList() ([]*model.Category, error) {
	var categories []*model.Category
	err := c.db.Debug().Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func NewCategoryDAO() *CategoryDAO {
	return &CategoryDAO{db: global.GetDB()}
}
