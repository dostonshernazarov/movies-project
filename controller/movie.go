package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dostonshernazarov/movies-app/middleware"
	"github.com/dostonshernazarov/movies-app/models"
	"github.com/dostonshernazarov/movies-app/services"
	"github.com/gin-gonic/gin"
)

// MovieController handles movie-related requests
type MovieController struct {
	MovieService *services.MovieService
}

// NewMovieController creates a new movie controller
func NewMovieController(movieService *services.MovieService) *MovieController {
	return &MovieController{
		MovieService: movieService,
	}
}

// @Summary Get all movies
// @Security BearerAuth
// @Description Get all movies from the database
// @Accept json
// @Produce json
// @Tags Movies
// @Success 200 {array} models.Movies
// @Failure 500 {object} models.ErrorResponse
// @Router /api/movies [get]
func (c *MovieController) GetAllMovies(ctx *gin.Context) {
	movies, err := c.MovieService.GetAllMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	movieResponses := make([]models.MovieResponse, len(movies))
	for i, movie := range movies {
		movieResponses[i] = models.MovieResponse{
			ID:        movie.ID,
			Title:     movie.Title,
			Director:  movie.Director,
			Year:      movie.Year,
			Plot:      movie.Plot,
			Genre:     movie.Genre,
			Rating:    movie.Rating,
			UserID:    movie.UserID,
			CreatedAt: movie.CreatedAt,
			UpdatedAt: movie.UpdatedAt,
		}
	}

	ctx.JSON(http.StatusOK, models.Movies{
		Movies: movieResponses,
	})
}

// @Summary Get a movie by ID
// @Security BearerAuth
// @Description Get a movie by ID from the database
// @Accept json
// @Produce json
// @Tags Movies
// @Param id path string true "Movie ID"
// @Success 200 {object} models.MovieResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/movies/{id} [get]
func (c *MovieController) GetMovieByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid ID format"})
		return
	}

	movie, err := c.MovieService.GetMovieByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Movie not found"})
		return
	}

	ctx.JSON(http.StatusOK, models.MovieResponse{
		ID:        uint(id),
		Title:     movie.Title,
		Director:  movie.Director,
		Year:      movie.Year,
		Plot:      movie.Plot,
		Genre:     movie.Genre,
		Rating:    movie.Rating,
		UserID:    movie.UserID,
		CreatedAt: movie.CreatedAt,
		UpdatedAt: movie.UpdatedAt,
	})
}

// @Summary Create a new movie
// @Security BearerAuth
// @Description Create a new movie with title, director, year, plot, genre, and rating
// @Accept json
// @Produce json
// @Tags Movies
// @Param movie body models.MovieRequest true "Movie details"
// @Success 201 {object} models.MovieCreateResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /api/movies [post]
func (c *MovieController) CreateMovie(ctx *gin.Context) {
	var movieRequest models.MovieRequest
	if err := ctx.ShouldBindJSON(&movieRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	// Get user ID from JWT token
	userID := middleware.GetUserID(ctx)

	movie := models.Movie{
		Title:    movieRequest.Title,
		Director: movieRequest.Director,
		Year:     movieRequest.Year,
		Plot:     movieRequest.Plot,
		Genre:    movieRequest.Genre,
		Rating:   movieRequest.Rating,
		UserID:   userID,
	}

	if err := c.MovieService.CreateMovie(&movie); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, models.MovieCreateResponse{
		Title:     movie.Title,
		Director:  movie.Director,
		Year:      movie.Year,
		Plot:      movie.Plot,
		Genre:     movie.Genre,
		Rating:    movie.Rating,
		CreatedAt: movie.CreatedAt.Format(time.RFC3339),
		UpdatedAt: movie.UpdatedAt.Format(time.RFC3339),
	})
}

// @Summary Update a movie
// @Security BearerAuth
// @Description Update a movie with title, director, year, plot, genre, and rating
// @Accept json
// @Produce json
// @Tags Movies
// @Param id path string true "Movie ID"
// @Param movie body models.MovieRequest true "Movie details"
// @Success 200 {object} models.MovieResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Router /api/movies/{id} [put]
func (c *MovieController) UpdateMovie(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid ID format"})
		return
	}

	var movieRequest models.MovieRequest
	if err := ctx.ShouldBindJSON(&movieRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	// Get existing movie
	existingMovie, err := c.MovieService.GetMovieByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Movie not found"})
		return
	}

	// Get user ID from JWT token
	userID := middleware.GetUserID(ctx)

	// Check if user owns the movie
	if existingMovie.UserID != userID {
		ctx.JSON(http.StatusForbidden, models.ErrorResponse{Error: "You don't have permission to update this movie"})
		return
	}

	// Update movie fields
	existingMovie.Title = movieRequest.Title
	existingMovie.Director = movieRequest.Director
	existingMovie.Year = movieRequest.Year
	existingMovie.Plot = movieRequest.Plot
	existingMovie.Genre = movieRequest.Genre
	existingMovie.Rating = movieRequest.Rating

	if err := c.MovieService.UpdateMovie(&existingMovie); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.MovieResponse{
		ID:        uint(id),
		Title:     existingMovie.Title,
		Director:  existingMovie.Director,
		Year:      existingMovie.Year,
		Plot:      existingMovie.Plot,
		Genre:     existingMovie.Genre,
		Rating:    existingMovie.Rating,
		UserID:    existingMovie.UserID,
		CreatedAt: existingMovie.CreatedAt,
		UpdatedAt: existingMovie.UpdatedAt,
	})
}

// @Summary Delete a movie
// @Security BearerAuth
// @Description Delete a movie by ID from the database
// @Accept json
// @Produce json
// @Tags Movies
// @Param id path string true "Movie ID"
// @Success 200 {object} string
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Router /api/movies/{id} [delete]
func (c *MovieController) DeleteMovie(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid ID format"})
		return
	}

	// Get existing movie
	existingMovie, err := c.MovieService.GetMovieByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Movie not found"})
		return
	}

	// Get user ID from JWT token
	userID := middleware.GetUserID(ctx)

	// Check if user owns the movie
	if existingMovie.UserID != userID {
		ctx.JSON(http.StatusForbidden, models.ErrorResponse{Error: "You don't have permission to delete this movie"})
		return
	}

	if err := c.MovieService.DeleteMovie(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
