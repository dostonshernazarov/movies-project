package controllers

import (
	"net/http"
	"strconv"

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

// GetAllMovies handles GET /api/movies
func (c *MovieController) GetAllMovies(ctx *gin.Context) {
	movies, err := c.MovieService.GetAllMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, movies)
}

// GetMovieByID handles GET /api/movies/:id
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

	ctx.JSON(http.StatusOK, movie)
}

// CreateMovie handles POST /api/movies
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

	ctx.JSON(http.StatusCreated, movie)
}

// UpdateMovie handles PUT /api/movies/:id
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

	ctx.JSON(http.StatusOK, existingMovie)
}

// DeleteMovie handles DELETE /api/movies/:id
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
