package repositories

import (
	"github.com/dostonshernazarov/movies-app/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByUsername(username string) (models.User, error) {
	var user models.User
	result := r.DB.Where("username = ?", username).First(&user)
	return user, result.Error
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}
