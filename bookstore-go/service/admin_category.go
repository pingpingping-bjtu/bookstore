package service

import (
	"bookstore-manager/model"
	"bookstore-manager/repository"
)

type AdminCategoryService struct {
	AdminCategoryDAO *repository.AdminCategoryDAO
}

func (g *AdminCategoryService) GetAdminCategories() (*[]model.Category, error) {
	return g.AdminCategoryDAO.GetAdminCategories()
}

func (g *AdminCategoryService) UpdateCategories(id uint, update map[string]interface{}) error {
	category, err := g.GetCategoryByID(id)
	if err != nil {
		return err
	}
	if name, ok := update["name"].(string); ok && name != "" {
		category.Name = name
	}
	if description, ok := update["description"].(string); ok {
		category.Description = description
	}
	if icon, ok := update["icon"].(string); ok {
		category.Icon = icon
	}
	if color, ok := update["color"].(string); ok {
		category.Color = color
	}
	if gradient, ok := update["gradient"].(string); ok {
		category.Gradient = gradient
	}
	if sort, ok := update["sort"].(int); ok {
		category.Sort = sort
	}
	if isActive, ok := update["is_active"].(bool); ok {
		category.IsActive = isActive
	}
	err = g.AdminCategoryDAO.UpdateCategories(category)
	if err != nil {
		return err
	}
	return nil
}

func (g *AdminCategoryService) GetCategoryByID(id uint) (*model.Category, error) {
	return g.AdminCategoryDAO.GetCategoryByID(id)
}

func (g *AdminCategoryService) CreateAdminCategory(category *model.Category) error {
	return g.AdminCategoryDAO.CreateAdminCategory(category)
}

func (g *AdminCategoryService) DeleteAdminCategory(id uint) error {
	return g.AdminCategoryDAO.DeleteAdminCategory(id)
}

func NewAdminCategoryService() *AdminCategoryService {
	return &AdminCategoryService{AdminCategoryDAO: repository.NewAdminCategoryDAO()}
}
