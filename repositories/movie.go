package repositories

import (
	"errors"

	"github.com/dostonshernazarov/movies-app/models"
	"gorm.io/gorm"
)

type MovieRepository struct {
	DB *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{DB: db}
}

func (r *MovieRepository) GetAll() ([]models.Movie, error) {
	var movies []models.Movie
	result := r.DB.Select("id, title, director, year, plot, genre, rating, user_id, created_at, updated_at").Find(&movies)
	return movies, result.Error
}

func (r *MovieRepository) GetByID(id uint) (models.Movie, error) {
	var movie models.Movie
	result := r.DB.First(&movie, id)
	return movie, result.Error
}

func (r *MovieRepository) Create(movie *models.Movie) error {
	return r.DB.Create(movie).Error
}

func (r *MovieRepository) Update(movie *models.Movie) error {
	result := r.DB.Save(movie)
	if result.RowsAffected == 0 {
		return errors.New("movie not found")
	}
	return result.Error
}

func (r *MovieRepository) Delete(id uint) error {
	result := r.DB.Delete(&models.Movie{}, id)
	if result.RowsAffected == 0 {
		return errors.New("movie not found")
	}
	return result.Error
}

func (r *MovieRepository) FindByUserID(userID uint) ([]models.Movie, error) {
	var movies []models.Movie
	result := r.DB.Where("user_id = ?", userID).Find(&movies)
	return movies, result.Error
}
