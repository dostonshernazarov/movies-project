package controllers

import (
	"net/http"

	"github.com/dostonshernazarov/movies-app/models"
	"github.com/dostonshernazarov/movies-app/services"
	"github.com/gin-gonic/gin"
)

// AuthController handles authentication-related requests
type AuthController struct {
	AuthService *services.AuthService
}

// NewAuthController creates a new authentication controller
func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

// Register handles POST /auth/register
func (c *AuthController) Register(ctx *gin.Context) {
	var request models.UserRegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	user := models.User{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
	}

	if err := c.AuthService.Register(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

// Login handles POST /auth/login
func (c *AuthController) Login(ctx *gin.Context) {
	var request models.UserLoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	token, user, err := c.AuthService.Login(request.Username, request.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.AuthResponse{
		Token: token,
		User:  user,
	})
}
