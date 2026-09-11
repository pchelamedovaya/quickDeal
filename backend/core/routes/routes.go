package routes

import (
	"quickdeal/controller"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, authController *controller.AuthController) {
	auth := router.Group("/api/auth")
	{
		auth.POST("/register", authController.Register)
	}
}
