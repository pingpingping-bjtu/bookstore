package controller

import (
	"bookstore-manager/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminDashboardController struct {
	AdminDashboardService *service.AdminDashboardService
}

func (d *AdminDashboardController) GetDashboardStats(c *gin.Context) {
	stats, err := d.AdminDashboardService.GetDashboardStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "获取统计数据成功",
		"data":    stats,
	})
}

func NewAdminDashboardController() *AdminDashboardController {
	return &AdminDashboardController{AdminDashboardService: service.NewAdminDashboardService()}
}
