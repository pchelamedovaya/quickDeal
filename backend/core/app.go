package core

import (
	"quickdeal/config"
	"quickdeal/controller"
	"quickdeal/core/routes"
	"quickdeal/repository"
	"quickdeal/service"

	"github.com/gin-gonic/gin"
)

func NewApp(cfg *config.AppConfig) (*gin.Engine, error) {
	db, err := config.ConnectDatabase(cfg.DBDsn)
	if err != nil {
		return nil, err
	}

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)

	router := gin.Default()
	routes.Register(router, authController)

	return router, nil
}
