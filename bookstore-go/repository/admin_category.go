package repository

import (
	"bookstore-manager/global"
	"bookstore-manager/model"
	"errors"

	"gorm.io/gorm"
)

type AdminCategoryDAO struct {
	db *gorm.DB
}

func (g *AdminCategoryDAO) GetAdminCategories() (*[]model.Category, error) {
	var categories *[]model.Category
	err := g.db.Debug().Model(&model.Category{}).Find(&categories).Error
	if err != nil {
		return nil, errors.New("获取分类列表失败")
	}
	return categories, nil

}

func (g *AdminCategoryDAO) GetCategoryByID(id uint) (*model.Category, error) {
	var category *model.Category
	if err := g.db.Debug().Where("id=?", id).Find(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (g *AdminCategoryDAO) UpdateCategories(update *model.Category) error {
	return g.db.Debug().Save(update).Error
}

func (g *AdminCategoryDAO) CreateAdminCategory(category *model.Category) error {
	return g.db.Debug().Save(category).Error
}

func (g *AdminCategoryDAO) DeleteAdminCategory(id uint) error {
	return g.db.Debug().Delete(&model.Category{}, id).Error

}

func NewAdminCategoryDAO() *AdminCategoryDAO {
	return &AdminCategoryDAO{db: global.GetDB()}
}
