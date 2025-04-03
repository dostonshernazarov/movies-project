package middleware

import (
	"net/http"
	"strings"

	"github.com/dostonshernazarov/movies-app/models"
	"github.com/dostonshernazarov/movies-app/services"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	jwtService := services.NewJWTService()

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Authorization header is required"})
			ctx.Abort()
			return
		}

		if strings.HasPrefix(authHeader, "Bearer ") {
			authHeader = strings.TrimPrefix(authHeader, "Bearer ")
		}

		token, err := jwtService.ValidateToken(authHeader)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid or expired token"})
			ctx.Abort()
			return
		}

		userID := jwtService.ExtractUserID(token)
		ctx.Set("user_id", userID)

		ctx.Next()
	}
}

func GetUserID(ctx *gin.Context) uint {
	userID, exists := ctx.Get("user_id")
	if !exists {
		return 0
	}
	return userID.(uint)
}
