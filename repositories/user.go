package repositories

import (
	"github.com/dostonshernazarov/movies-app/models"
	"gorm.io/gorm"
)

// UserRepository handles database operations for users
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// FindByUsername finds a user by username
func (r *UserRepository) FindByUsername(username string) (models.User, error) {
	var user models.User
	result := r.DB.Where("username = ?", username).First(&user)
	return user, result.Error
}

// Create creates a new user
func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}
