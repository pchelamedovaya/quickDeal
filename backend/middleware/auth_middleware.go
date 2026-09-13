package middleware

import (
	"net/http"
	"strings"

	"quickdeal/utils"

	"github.com/gin-gonic/gin"
)

const UserIDContextKey = "user_id"

func AuthMiddleware(accessSecret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		token, ok := strings.CutPrefix(header, "Bearer ")
		if !ok || token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid authorization header"})
			return
		}

		userID, err := utils.ParseAccessToken(token, accessSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired access token"})
			return
		}

		ctx.Set(UserIDContextKey, userID)
		ctx.Next()
	}
}
