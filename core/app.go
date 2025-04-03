package core

import (
	"github.com/dostonshernazarov/movies-app/config"
	controllers "github.com/dostonshernazarov/movies-app/controller"
	_ "github.com/dostonshernazarov/movies-app/docs"
	"github.com/dostonshernazarov/movies-app/middleware"
	"github.com/dostonshernazarov/movies-app/repositories"
	"github.com/dostonshernazarov/movies-app/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// App represents the application
type App struct {
	Engine *gin.Engine
}

// Run starts the application server
func (a *App) Run(addr string) error {
	return a.Engine.Run(addr)
}

// InitializeApp initializes and returns the application
func InitializeApp() (*App, error) {
	var app App

	err := fx.New(
		// Provide database connection
		fx.Provide(config.NewDatabaseConnection),

		// Provide repositories
		fx.Provide(repositories.NewMovieRepository),
		fx.Provide(repositories.NewUserRepository),

		// Provide services
		fx.Provide(services.NewJWTService),
		fx.Provide(services.NewMovieService),
		fx.Provide(services.NewAuthService),

		// Provide controllers
		fx.Provide(controllers.NewAuthController),
		fx.Provide(controllers.NewMovieController),

		// Provide Gin engine with routes
		fx.Provide(NewGinEngine),

		// Inject the engine into the app
		fx.Populate(&app.Engine),
	).Err()

	return &app, err
}

// NewGinEngine creates and configures the Gin engine with routes
// @title Movies API
// @version 1.0
// @description API for managing movies
// @host localhost:8060
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @type string
// @description The token for the user
// @required
// @default Bearer
func NewGinEngine(
	movieController *controllers.MovieController,
	authController *controllers.AuthController,
) *gin.Engine {
	engine := gin.Default()

	// Add middleware
	engine.Use(middleware.CORSMiddleware())

	// Auth routes
	authRoutes := engine.Group("/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
	}

	// API routes with JWT auth
	apiRoutes := engine.Group("/api")
	apiRoutes.Use(middleware.JWTAuthMiddleware())
	{
		// Movie routes
		movies := apiRoutes.Group("/movies")
		{
			movies.GET("", movieController.GetAllMovies)
			movies.GET("/:id", movieController.GetMovieByID)
			movies.POST("", movieController.CreateMovie)
			movies.PUT("/:id", movieController.UpdateMovie)
			movies.DELETE("/:id", movieController.DeleteMovie)
		}
	}

	url := ginSwagger.URL("swagger/doc.json")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return engine
}
