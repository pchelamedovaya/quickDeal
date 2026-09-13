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
	refreshTokenRepo := repository.NewRefreshTokenRepository(db)
	authService := service.NewAuthService(
		userRepo,
		refreshTokenRepo,
		cfg.JWTAccessSecret,
		cfg.AccessTokenTTL,
		cfg.RefreshTokenTTL,
	)
	authController := controller.NewAuthController(authService)

	router := gin.Default()
	routes.Register(router, authController)

	return router, nil
}
