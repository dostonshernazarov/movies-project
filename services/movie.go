package services

import (
	"github.com/dostonshernazarov/movies-app/models"
	"github.com/dostonshernazarov/movies-app/repositories"
	"gorm.io/gorm"
)

type MovieService struct {
	MovieRepo *repositories.MovieRepository
	DB        *gorm.DB
}

func NewMovieService(movieRepo *repositories.MovieRepository, db *gorm.DB) *MovieService {
	return &MovieService{
		MovieRepo: movieRepo,
		DB:        db,
	}
}

func (s *MovieService) GetAllMovies() ([]models.Movie, error) {
	return s.MovieRepo.GetAll()
}

func (s *MovieService) GetMovieByID(id uint) (models.Movie, error) {
	return s.MovieRepo.GetByID(id)
}

func (s *MovieService) CreateMovie(movie *models.Movie) error {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		return s.MovieRepo.Create(movie)
	})
}

func (s *MovieService) UpdateMovie(movie *models.Movie) error {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		return s.MovieRepo.Update(movie)
	})
}

func (s *MovieService) DeleteMovie(id uint) error {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		return s.MovieRepo.Delete(id)
	})
}

func (s *MovieService) GetUserMovies(userID uint) ([]models.Movie, error) {
	return s.MovieRepo.FindByUserID(userID)
}
