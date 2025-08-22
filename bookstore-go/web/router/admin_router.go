package router

import (
	"bookstore-manager/web/controller"

	"github.com/gin-gonic/gin"
)

func InitAdminRouter() *gin.Engine {
	r := gin.Default()

	// 添加CORS中间件
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	adminUserController := controller.NewAdminUserController()
	v1 := r.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.POST("/auth/login", adminUserController.AdminUserLogin)
		}
	}
	return r
}
