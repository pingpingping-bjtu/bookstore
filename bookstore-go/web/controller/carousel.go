package controller

import (
	"bookstore-manager/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CarouselController struct {
	CarouselService *service.CarouselService
}

func (l *CarouselController) GetCarouselList(c *gin.Context) {
	carousels, err := l.CarouselService.GetCarouselList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "获取轮播图失败",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "获取轮播图成功",
		"data":    carousels,
	})
}

func NewCarouselController() *CarouselController {
	return &CarouselController{CarouselService: service.NewCarouselService()}
}
