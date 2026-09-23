package controller

import (
	"errors"
	"net/http"

	"quickdeal/apperrors"
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
		apperrors.Validation(ctx, err)
		return
	}

	user, err := c.authService.Register(req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrEmailTaken):
			apperrors.JSON(ctx, http.StatusConflict, apperrors.CodeEmailTaken)
		case errors.Is(err, service.ErrUsernameTaken):
			apperrors.JSON(ctx, http.StatusConflict, apperrors.CodeUsernameTaken)
		default:
			apperrors.Internal(ctx)
		}
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apperrors.Validation(ctx, err)
		return
	}

	tokens, err := c.authService.Login(req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidCredentials):
			apperrors.JSON(ctx, http.StatusUnauthorized, apperrors.CodeInvalidCredentials)
		default:
			apperrors.Internal(ctx)
		}
		return
	}

	ctx.JSON(http.StatusOK, tokens)
}

func (c *AuthController) Refresh(ctx *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apperrors.Validation(ctx, err)
		return
	}

	tokens, err := c.authService.Refresh(req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidRefreshToken):
			apperrors.JSON(ctx, http.StatusUnauthorized, apperrors.CodeInvalidRefreshToken)
		default:
			apperrors.Internal(ctx)
		}
		return
	}

	ctx.JSON(http.StatusOK, tokens)
}

func (c *AuthController) Logout(ctx *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apperrors.Validation(ctx, err)
		return
	}

	if err := c.authService.Logout(req); err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidRefreshToken):
			apperrors.JSON(ctx, http.StatusUnauthorized, apperrors.CodeInvalidRefreshToken)
		default:
			apperrors.Internal(ctx)
		}
		return
	}

	ctx.Status(http.StatusNoContent)
}
