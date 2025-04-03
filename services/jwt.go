package services

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dostonshernazarov/movies-app/models"
	"github.com/golang-jwt/jwt/v4"
)

// JWTService handles JWT operations
type JWTService struct {
	secretKey string
	issuer    string
}

// NewJWTService creates a new JWT service
func NewJWTService() *JWTService {
	return &JWTService{
		secretKey: getEnv("JWT_SECRET", "your-secret-key"),
		issuer:    "movies-api",
	}
}

// GenerateToken generates a new JWT token
func (s *JWTService) GenerateToken(user models.User) string {
	// Set expiration time (24 hours)
	expDuration, _ := strconv.Atoi(getEnv("TOKEN_HOUR_LIFESPAN", "24"))
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"iss":      s.issuer,
		"exp":      time.Now().Add(time.Hour * time.Duration(expDuration)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		fmt.Println("Error signing token:", err)
		return ""
	}

	return tokenString
}

// ValidateToken validates a JWT token
func (s *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
}

// ExtractUserID extracts the user ID from a JWT token
func (s *JWTService) ExtractUserID(token *jwt.Token) uint {
	claims := token.Claims.(jwt.MapClaims)
	id, _ := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
	return uint(id)
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
