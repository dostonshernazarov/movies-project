package repositories

import (
	"errors"

	"github.com/dostonshernazarov/movies-app/models"
	"gorm.io/gorm"
)

// MovieRepository handles database operations for movies
type MovieRepository struct {
	DB *gorm.DB
}

// NewMovieRepository creates a new movie repository
func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{DB: db}
}

// GetAll retrieves all movies
func (r *MovieRepository) GetAll() ([]models.Movie, error) {
	var movies []models.Movie
	result := r.DB.Find(&movies)
	return movies, result.Error
}

// GetByID retrieves a movie by ID
func (r *MovieRepository) GetByID(id uint) (models.Movie, error) {
	var movie models.Movie
	result := r.DB.First(&movie, id)
	return movie, result.Error
}

// Create creates a new movie
func (r *MovieRepository) Create(movie *models.Movie) error {
	return r.DB.Create(movie).Error
}

// Update updates an existing movie
func (r *MovieRepository) Update(movie *models.Movie) error {
	result := r.DB.Save(movie)
	if result.RowsAffected == 0 {
		return errors.New("movie not found")
	}
	return result.Error
}

// Delete deletes a movie by ID
func (r *MovieRepository) Delete(id uint) error {
	result := r.DB.Delete(&models.Movie{}, id)
	if result.RowsAffected == 0 {
		return errors.New("movie not found")
	}
	return result.Error
}

// FindByUserID finds all movies created by a user
func (r *MovieRepository) FindByUserID(userID uint) ([]models.Movie, error) {
	var movies []models.Movie
	result := r.DB.Where("user_id = ?", userID).Find(&movies)
	return movies, result.Error
}
