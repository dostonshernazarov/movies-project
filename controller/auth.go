package controllers

import (
	"net/http"
	"time"

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

// @Summary Register a new user
// @Description Register a new user with username, password, and email
// @Accept json
// @Produce json
// @Tags Auth
// @Param user body models.UserRegisterRequest true "User registration details"
// @Success 201 {object} models.UserRegisterResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/register [post]
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

	ctx.JSON(http.StatusCreated, models.UserRegisterResponse{
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	})
}

// @Summary Login a user
// @Description Login a user with username and password
// @Accept json
// @Produce json
// @Tags Auth
// @Param user body models.UserLoginRequest true "User login details"
// @Success 200 {object} models.AuthResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /auth/login [post]
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
		Token:     token,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	})
}
