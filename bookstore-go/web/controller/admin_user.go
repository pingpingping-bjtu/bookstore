package controller

import (
	"bookstore-manager/repository"
	"bookstore-manager/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminUserController struct {
	AdminUserService *service.AdminUserService
}

type AdminLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *AdminUserController) AdminUserLogin(c *gin.Context) {
	var req AdminLoginRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}
	response, err := u.AdminUserService.AdminUserLogin(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"data":    response,
		"message": "登录成功",
	})

}

func (u *AdminUserController) GetUsersList(c *gin.Context) {

	req := &repository.GetUsersRequest{}
	// 获取查询参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	req.Username = c.Query("username")
	req.Email = c.Query("email")
	req.IsAdmin = c.Query("is_admin")
	users, total, err := u.AdminUserService.GetUsersList(req, page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "查询用户失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "查询用户成功",
		"data": gin.H{
			"users":        users,
			"total":        total,
			"current_page": page,
			"page_size":    pageSize,
		},
	})

}

func (u *AdminUserController) UpdateUserStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "ID参数错误",
		})
		return
	}

	var req struct {
		IsAdmin bool `json:"is_admin"`
	}
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "参数绑定错误",
		})
		return
	}
	err = u.AdminUserService.UpdateUserStatus(req.IsAdmin, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "更新状态失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "更新状态成功",
	})
}

func (u *AdminUserController) CreateUser(c *gin.Context) {
	req := &repository.CreateUserRequest{}
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "参数绑定失败",
		})
		return
	}
	user, err := u.AdminUserService.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "用户创建失败",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "用户创建成功",
		"data":    user,
	})
}

func (u *AdminUserController) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "无效id",
		})
		return
	}
	err = u.AdminUserService.DeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "删除用户失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "删除用户成功",
	})
	return

}

func (u *AdminUserController) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "无效id",
		})
		return
	}
	req := &repository.UpdateUserRequest{}
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "参数绑定失败",
		})
		return
	}
	user, err := u.AdminUserService.UpdateUser(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "更新用户信息失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "更新用户信息成功",
		"data":    user,
	})

}

func NewAdminUserController() *AdminUserController {
	return &AdminUserController{
		AdminUserService: service.NewAdminUserService(),
	}

}
