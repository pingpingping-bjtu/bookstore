package router

import (
	"bookstore-manager/web/controller"
	"bookstore-manager/web/middleware"

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
	adminDashboardController := controller.NewAdminDashboardController()
	adminCategoriesController := controller.NewAdminCategoryController()
	adminBookController := controller.NewAdminBookController()
	v1 := r.Group("/api/v1")
	{
		login := v1.Group("/admin/auth")
		{
			login.POST("/login", adminUserController.AdminUserLogin)
		}
		admin := v1.Group("/admin")
		admin.Use(middleware.JWTAuthMiddleware())
		{
			admin.GET("/dashboard/stats", adminDashboardController.GetDashboardStats)

			//分类路由
			categories := admin.Group("/categories")
			{
				categories.GET("/list", adminCategoriesController.GetAdminCategories)
				categories.POST("/create", adminCategoriesController.CreateAdminCategories)
				categories.PUT("/:id", adminCategoriesController.UpdateAdminCategories)
				categories.DELETE("/:id", adminCategoriesController.DeleteAdminCategory)
			}
			//用户管理路由
			users := admin.Group("/users")
			{
				users.GET("/list", adminUserController.GetUsersList)
				users.PUT("/:id/status", adminUserController.UpdateUserStatus)
				users.POST("/create", adminUserController.CreateUser)
				users.DELETE("/:id", adminUserController.DeleteUser)
				users.PUT("/:id", adminUserController.UpdateUser)
			}
			//图书管理
			books := admin.Group("/books")
			{
				books.GET("/list", adminBookController.GetBookList)
				books.GET("/:id", adminBookController.GetBookByID)
				books.POST("/create", adminBookController.CreateBook)
				books.PUT("/:id", adminBookController.UpdateBook)
				books.DELETE("/:id", adminBookController.DeleteBook)
				books.PUT("/:id/status", adminBookController.UpdateBookStatus)
			}

		}

	}
	return r
}
