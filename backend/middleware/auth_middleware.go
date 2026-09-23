package middleware

import (
	"net/http"
	"strings"

	"quickdeal/apperrors"
	"quickdeal/utils"

	"github.com/gin-gonic/gin"
)

const UserIDContextKey = "user_id"

func AuthMiddleware(accessSecret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		token, ok := strings.CutPrefix(header, "Bearer ")
		if !ok || token == "" {
			apperrors.AbortJSON(ctx, http.StatusUnauthorized, apperrors.CodeMissingAuthHeader)
			return
		}

		userID, err := utils.ParseAccessToken(token, accessSecret)
		if err != nil {
			apperrors.AbortJSON(ctx, http.StatusUnauthorized, apperrors.CodeInvalidOrExpiredToken)
			return
		}

		ctx.Set(UserIDContextKey, userID)
		ctx.Next()
	}
}
