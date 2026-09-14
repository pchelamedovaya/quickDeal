package routes

import (
	"quickdeal/controller"
	"quickdeal/middleware"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, authController *controller.AuthController, adController *controller.AdController, accessSecret string) {
	auth := router.Group("/api/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
		auth.POST("/refresh", authController.Refresh)
		auth.POST("/logout", authController.Logout)
	}

	ads := router.Group("/api/ads")
	ads.Use(middleware.AuthMiddleware(accessSecret))
	{
		ads.POST("", adController.Create)
	}
}
