package apperrors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CodeEmailTaken            = "EMAIL_TAKEN"
	CodeUsernameTaken         = "USERNAME_TAKEN"
	CodeInvalidCredentials    = "INVALID_CREDENTIALS"
	CodeInvalidRefreshToken   = "INVALID_REFRESH_TOKEN"
	CodeMissingAuthHeader     = "MISSING_AUTH_HEADER"
	CodeInvalidOrExpiredToken = "INVALID_OR_EXPIRED_TOKEN"
	CodeValidationError       = "VALIDATION_ERROR"
	CodeInternalError         = "INTERNAL_ERROR"
)

type Response struct {
	ErrorCode string `json:"error_code"`
	Details   string `json:"details,omitempty"`
}

func JSON(ctx *gin.Context, status int, code string) {
	ctx.JSON(status, Response{ErrorCode: code})
}

func AbortJSON(ctx *gin.Context, status int, code string) {
	ctx.AbortWithStatusJSON(status, Response{ErrorCode: code})
}

func Validation(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, Response{
		ErrorCode: CodeValidationError,
		Details:   err.Error(),
	})
}

func Internal(ctx *gin.Context) {
	JSON(ctx, http.StatusInternalServerError, CodeInternalError)
}
