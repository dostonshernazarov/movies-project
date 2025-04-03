package config

import (
	"fmt"
	"log"
	"os"

	"github.com/dostonshernazarov/movies-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDatabaseConnection creates a new database connection
func NewDatabaseConnection() *gorm.DB {
	dbConfig := DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "doston"),
		DBName:   getEnv("DB_NAME", "movies_db"),
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate tables
	if err := db.AutoMigrate(&models.User{}, &models.Movie{}); err != nil {
		log.Fatalf("Failed to auto migrate tables: %v", err)
	}

	return db
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
