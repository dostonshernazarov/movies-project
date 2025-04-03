package services

import (
	"errors"

	"github.com/dostonshernazarov/movies-app/models"
	"github.com/dostonshernazarov/movies-app/repositories"
	"golang.org/x/crypto/bcrypt"
)

// AuthService handles authentication business logic
type AuthService struct {
	UserRepo   *repositories.UserRepository
	JWTService *JWTService
}

// NewAuthService creates a new authentication service
func NewAuthService(userRepo *repositories.UserRepository, jwtService *JWTService) *AuthService {
	return &AuthService{
		UserRepo:   userRepo,
		JWTService: jwtService,
	}
}

// Register registers a new user
func (s *AuthService) Register(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.UserRepo.Create(user)
}

// Login authenticates a user
func (s *AuthService) Login(username, password string) (string, models.User, error) {
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		return "", models.User{}, errors.New("user not found")
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", models.User{}, errors.New("invalid credentials")
	}

	// Generate JWT token
	token := s.JWTService.GenerateToken(user)
	return token, user, nil
}
