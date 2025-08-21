package router

import (
	"bookstore-manager/web/controller"
	"bookstore-manager/web/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
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

	userController := controller.NewUserController()
	captchaController := controller.NewCaptchaController()
	bookController := controller.NewBookController()
	favoriteController := controller.NewFavoriteController()
	orderController := controller.NewOrderController()
	categoryController := controller.NewCategoryController()
	carouselController := controller.NewCarouselController()
	v1 := r.Group("api/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/register", userController.UserRegister)
			user.POST("/login", userController.UserLogin)
		}
		auth := user.Group("")
		{
			auth.Use(middleware.JWTAuthMiddleware())
			{
				auth.GET("/profile", userController.GetUserProfile)
				auth.PUT("/profile", userController.UpdateUserProfile)
				auth.PUT("/password", userController.ChangePassword)
				auth.DELETE("/logout", userController.LogOut)
			}
		}
		book := v1.Group("/book")
		{
			book.GET("/hot", bookController.GetHotBooks)
			book.GET("/new", bookController.GetNewBooks)
			book.GET("/list", bookController.GetBookList)
			book.GET("/search", bookController.SearchBook)
			book.GET("/detail/:id", bookController.GetBookDetail)
			book.GET("/category/:category", bookController.GetBookByCategory)
		}
		favorite := v1.Group("/favorite")
		favorite.Use(middleware.JWTAuthMiddleware())
		{
			favorite.POST("/:id", favoriteController.AddFavoriteBook)
			favorite.DELETE("/:id", favoriteController.RemoveFavoriteBook)
			favorite.GET("/list", favoriteController.GetFavoriteList)
			favorite.GET("/:id/check", favoriteController.CheckFavorite)
			favorite.GET("/count", favoriteController.GetFavoriteCount)
		}
		order := v1.Group("/order")
		order.Use(middleware.JWTAuthMiddleware())
		{
			order.POST("/create", orderController.CreateOrder)
			order.GET("/list", orderController.GetOrderList)
			order.POST("/:id/pay", orderController.PayOrder)
		}
		category := v1.Group("/category")
		{
			category.GET("/list", categoryController.GetCategoryList)
		}
		carousel := v1.Group("/carousel")
		{
			carousel.GET("/list", carouselController.GetCarouselList)
		}
	}
	//验证图形验证码
	captcha := v1.Group("/captcha")
	{
		captcha.GET("/generate", captchaController.GenerateCaptcha)
	}
	return r
}
