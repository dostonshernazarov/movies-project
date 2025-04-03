package middleware

import (
	"net/http"
	"strings"

	"github.com/dostonshernazarov/movies-app/models"
	"github.com/dostonshernazarov/movies-app/services"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware is a middleware that validates JWT tokens
func JWTAuthMiddleware() gin.HandlerFunc {
	jwtService := services.NewJWTService()

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Authorization header is required"})
			ctx.Abort()
			return
		}

		// If authorization header has Bearer prefix, remove it
		if strings.HasPrefix(authHeader, "Bearer ") {
			authHeader = strings.TrimPrefix(authHeader, "Bearer ")
		}

		// Extract token

		// Validate token
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid or expired token"})
			ctx.Abort()
			return
		}

		// Extract user ID from token
		userID := jwtService.ExtractUserID(token)
		ctx.Set("user_id", userID)

		ctx.Next()
	}
}

// GetUserID gets the user ID from the context
func GetUserID(ctx *gin.Context) uint {
	userID, exists := ctx.Get("user_id")
	if !exists {
		return 0
	}
	return userID.(uint)
}
