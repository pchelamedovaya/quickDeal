package controller

import (
	"net/http"

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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := ctx.MustGet(middleware.UserIDContextKey).(uuid.UUID)

	ad, err := c.adService.Create(userID, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create ad"})
		return
	}

	ctx.JSON(http.StatusCreated, ad)
}
