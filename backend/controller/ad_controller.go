package controller

import (
	"net/http"

	"quickdeal/apperrors"
	"quickdeal/dto"
	"quickdeal/middleware"
	"quickdeal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AdController struct {
	adService *service.AdService
}

func NewAdController(adService *service.AdService) *AdController {
	return &AdController{adService: adService}
}

func (c *AdController) Create(ctx *gin.Context) {
	var req dto.CreateAdRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apperrors.Validation(ctx, err)
		return
	}

	userID := ctx.MustGet(middleware.UserIDContextKey).(uuid.UUID)

	ad, err := c.adService.Create(userID, req)
	if err != nil {
		apperrors.Internal(ctx)
		return
	}

	ctx.JSON(http.StatusCreated, ad)
}

func (c *AdController) List(ctx *gin.Context) {
	ads, err := c.adService.List()
	if err != nil {
		apperrors.Internal(ctx)
		return
	}

	ctx.JSON(http.StatusOK, ads)
}
