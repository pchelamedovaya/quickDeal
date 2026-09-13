package controller

import (
	"errors"
	"net/http"

	"quickdeal/dto"
	"quickdeal/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.authService.Register(req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrEmailTaken), errors.Is(err, service.ErrUsernameTaken):
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := c.authService.Login(req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidCredentials):
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log in"})
		}
		return
	}

	ctx.JSON(http.StatusOK, tokens)
}

func (c *AuthController) Refresh(ctx *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := c.authService.Refresh(req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidRefreshToken):
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to refresh token"})
		}
		return
	}

	ctx.JSON(http.StatusOK, tokens)
}

func (c *AuthController) Logout(ctx *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.authService.Logout(req); err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidRefreshToken):
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log out"})
		}
		return
	}

	ctx.Status(http.StatusNoContent)
}
