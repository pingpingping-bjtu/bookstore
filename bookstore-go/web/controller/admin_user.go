package controller

import (
	"bookstore-manager/service"
	"net/http"

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
		c.JSON(http.StatusInternalServerError, gin.H{
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

func NewAdminUserController() *AdminUserController {
	return &AdminUserController{
		AdminUserService: service.NewAdminUserService(),
	}

}
