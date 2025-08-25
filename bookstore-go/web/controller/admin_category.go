package controller

import (
	"bookstore-manager/model"
	"bookstore-manager/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminCategoryController struct {
	AdminCategoryService *service.AdminCategoryService
}

func (g *AdminCategoryController) GetAdminCategories(c *gin.Context) {
	categories, err := g.AdminCategoryService.GetAdminCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "获取分类成功",
		"data":    categories,
	})
}

func (g *AdminCategoryController) UpdateAdminCategories(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "无效的分类id",
		})
		return
	}
	var update map[string]interface{}
	err = c.ShouldBindBodyWithJSON(&update)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "参数绑定错误",
		})
		return
	}
	err = g.AdminCategoryService.UpdateCategories(uint(id), update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "分类更新失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "分类更新成功",
	})

}

func (g *AdminCategoryController) DeleteAdminCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "无效的分类id",
		})
		return
	}
	err = g.AdminCategoryService.DeleteAdminCategory(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "删除分类失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "删除分类成功",
	})
	return
}

func (g *AdminCategoryController) CreateAdminCategories(c *gin.Context) {
	var category *model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "参数绑定错误",
		})
		return
	}
	err := g.AdminCategoryService.CreateAdminCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "分类创建失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "分类创建成功",
	})

}

func NewAdminCategoryController() *AdminCategoryController {
	return &AdminCategoryController{AdminCategoryService: service.NewAdminCategoryService()}
}
