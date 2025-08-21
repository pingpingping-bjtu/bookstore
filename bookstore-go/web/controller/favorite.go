package controller

import (
	"bookstore-manager/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavoriteController struct {
	FavoriteService *service.FavoriteService
}

func NewFavoriteController() *FavoriteController {
	return &FavoriteController{FavoriteService: service.NewFavoriteService()}
}
func getUserID(c *gin.Context) int {
	userID, exists := c.Get("userID")
	if !exists {
		return 0
	}
	return userID.(int)
}
func (f *FavoriteController) AddFavoriteBook(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "请先登录",
		})
		return
	}
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "书籍ID不存在",
		})
		return
	}
	err = f.FavoriteService.AddFavorite(userID, bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "收藏失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "收藏成功",
	})
}

func (f *FavoriteController) RemoveFavoriteBook(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "请先登录",
		})
		return
	}
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "书籍ID不存在",
		})
		return
	}
	err = f.FavoriteService.RemoveFavorite(userID, bookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "取消收藏失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "取消收藏成功",
	})

}

func (f *FavoriteController) GetFavoriteList(c *gin.Context) {
	//获取用户信息
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    -1,
			"message": "请先登录",
		})
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "12"))
	//time_filter := c.DefaultQuery("time_filter", "all")
	favorites, total, err := f.FavoriteService.GetFavoriteList(userID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "收藏夹获取失败",
		})
		return
	}
	totalPages := (int(total) + pageSize - 1) / pageSize
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "收藏夹获取成功",
		"data": gin.H{
			"favorites":    favorites,
			"total":        total,
			"total_pages":  totalPages,
			"current_page": page,
		},
	})
}

func (f *FavoriteController) CheckFavorite(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    -1,
			"message": "请先登录",
		})
		return
	}
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "无效id",
		})
		return
	}
	isFavorite, err := f.FavoriteService.CheckFavorite(userID, bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "获取收藏信息失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"is_favorite": isFavorite,
		},
	})

}

func (f *FavoriteController) GetFavoriteCount(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    -1,
			"message": "请先登录",
		})
		return
	}
	count, err := f.FavoriteService.GetFavoriteCount(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "收藏数量获取失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "收藏数量获取成功",
		"data": gin.H{
			"count": count,
		},
	})

}
