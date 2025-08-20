package controller

import (
	"bookstore-manager/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookService *service.BookService
}

func NewBookController() *BookController {
	return &BookController{
		BookService: service.NewBookService(),
	}
}

func (b *BookController) GetHotBooks(c *gin.Context) {
	//根据销量降序排列
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	books, err := b.BookService.GetHotBooks(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "获取热销书失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"data":    books,
		"message": "获取热销书成功",
	})

}

func (b *BookController) GetNewBooks(c *gin.Context) {
	//根据上架时间排序
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	books, err := b.BookService.GetNewBooks(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "获取新书失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"data":    books,
		"message": "获取新书成功",
	})

}

func (b *BookController) GetBookList(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "12"))
	books, total, err := b.BookService.GetBookByPage(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "获取书籍列表失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"books":      books,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
			"total_size": (total + int64(pageSize-1)) / int64(pageSize),
		},
		"message": "获取书籍列表成功",
	})

}
