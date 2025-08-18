package router

import (
	"bookstore-manager/web/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	//r.GET("/test", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"data":  "hello",
	//		"error": "none",
	//	})
	//})

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
	v1 := r.Group("api/v1")
	{
		user := v1.Group("/usr")
		{
			user.POST("/register", controller.UserRegister)
			user.POST("/login", controller.UserLogin)
		}
	}
	return r
}
