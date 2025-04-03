package main

import (
	"log"
	"os"

	"github.com/dostonshernazarov/movies-app/config"
	"github.com/dostonshernazarov/movies-app/core"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	app, err := core.InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = config.DefaultPort
	}

	if err := app.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
