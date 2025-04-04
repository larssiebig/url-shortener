package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/larssiebig/url-shortener/internal/cache"
	"github.com/larssiebig/url-shortener/internal/handlers"
	"github.com/larssiebig/url-shortener/internal/repository"
)

func main() {
	// Load .env file from the root of the container
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file in main file")
	}

	// Initialize database and cache
	repository.InitDB()
	cache.InitRedis()

	// Create new Gin router
	r := gin.Default()

	// Define routes
	r.POST("/shorten", handlers.ShortenURL)
	r.GET("/:shortcode", handlers.RedirectURL)

	// Start the server
	log.Println("Starting server on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
