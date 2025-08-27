package controller

import (
	"bookstore-manager/model"
	"bookstore-manager/repository"
	"bookstore-manager/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminBookController struct {
	AdminBookService *service.AdminBookService
}

func (b *AdminBookController) GetBookList(c *gin.Context) {
	//分页
	req := &repository.GetBookRequest{}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "12"))
	req.Title = c.Query("title")
	req.Author = c.Query("author")
	req.BookType = c.Query("type")
	req.Status = c.Query("status")

	books, total, err := b.AdminBookService.GetBookList(page, pageSize, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "获取书籍列表失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "获取图书列表成功",
		"data": gin.H{
			"books":      books,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
			"total_size": (total + int64(pageSize-1)) / int64(pageSize),
		},
	})
}

func (b *AdminBookController) CreateBook(c *gin.Context) {
	req := model.BookCreateRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "参数绑定失败！",
		})
		return
	}
	err = b.AdminBookService.CreateBook(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "新增书本失败！",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "新增书本成功！",
	})
}

func (b *AdminBookController) GetBookByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "无效id！",
		})
		return
	}
	book, err := b.AdminBookService.GetBookByID(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "获取图书信息失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "获取图书信息成功",
		"data":    book,
	})
	return
}

func (b *AdminBookController) UpdateBookStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "ID参数错误",
		})
		return
	}
	statusStr := c.Query("status")
	if statusStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "缺少查询参数",
		})
		return
	}
	status, err := strconv.Atoi(statusStr)
	if err != nil || (status != 0 && status != 1) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "状态参数错误，只能是0或1",
		})
		return
	}
	err = b.AdminBookService.UpdateBookStatus(uint(id), status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "更新图书状态失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "更新图书状态成功",
	})
}

func (b *AdminBookController) DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "ID参数错误",
		})
		return
	}
	err = b.AdminBookService.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "删除书本失败！",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "删除书本成功！",
	})
}

func (b *AdminBookController) UpdateBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "ID参数错误",
		})
		return
	}

	var req model.BookUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "参数错误: " + err.Error(),
		})
		return
	}
	err = b.AdminBookService.UpdateBook(int(id), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "更新图书失败: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "更新图书成功",
	})

}

func NewAdminBookController() *AdminBookController {
	return &AdminBookController{AdminBookService: service.NewAdminBookService()}
}
