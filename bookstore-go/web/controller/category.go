package controller

import (
	"bookstore-manager/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	CategoryService *service.CategoryService
}

// GetCategoryList 获取所有分类
func (l *CategoryController) GetCategoryList(c *gin.Context) {
	categories, err := l.CategoryService.GetCategoryList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "获取分类列表失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"data":    categories,
		"message": "获取分类列表成功",
	})
}

func NewCategoryController() *CategoryController {
	return &CategoryController{CategoryService: service.NewCategoryService()}
}
