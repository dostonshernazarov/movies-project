package services

import (
	"github.com/dostonshernazarov/movies-app/models"
	"github.com/dostonshernazarov/movies-app/repositories"
	"gorm.io/gorm"
)

// MovieService handles business logic for movies
type MovieService struct {
	MovieRepo *repositories.MovieRepository
	DB        *gorm.DB
}

// NewMovieService creates a new movie service
func NewMovieService(movieRepo *repositories.MovieRepository, db *gorm.DB) *MovieService {
	return &MovieService{
		MovieRepo: movieRepo,
		DB:        db,
	}
}

// GetAllMovies retrieves all movies
func (s *MovieService) GetAllMovies() ([]models.Movie, error) {
	return s.MovieRepo.GetAll()
}

// GetMovieByID retrieves a movie by ID
func (s *MovieService) GetMovieByID(id uint) (models.Movie, error) {
	return s.MovieRepo.GetByID(id)
}

// CreateMovie creates a new movie
func (s *MovieService) CreateMovie(movie *models.Movie) error {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		// Using transaction in case we need to perform related operations
		return s.MovieRepo.Create(movie)
	})
}

// UpdateMovie updates an existing movie
func (s *MovieService) UpdateMovie(movie *models.Movie) error {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		return s.MovieRepo.Update(movie)
	})
}

// DeleteMovie deletes a movie by ID
func (s *MovieService) DeleteMovie(id uint) error {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		return s.MovieRepo.Delete(id)
	})
}

// GetUserMovies retrieves all movies created by a user
func (s *MovieService) GetUserMovies(userID uint) ([]models.Movie, error) {
	return s.MovieRepo.FindByUserID(userID)
}
