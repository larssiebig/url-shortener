package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Load .env file
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve the database connection string from environment variables
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set in .env")
	}

	// Initialize database connection with retries
	var db *sql.DB
	for retries := 0; retries < 5; retries++ {
		db, err = sql.Open("postgres", dbURL)
		if err != nil {
			log.Printf("Failed to connect to database, retrying... (%d/5)", retries+1)
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	if err != nil {
		log.Fatal("Database connection failed after retries: ", err)
	}
	defer db.Close()

	// Check if the database is reachable
	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	// Create new Gin router
	r := gin.Default()

	// Define routes
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// Start the server
	log.Println("Starting server on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
