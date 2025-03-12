package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

func InitDB() {
	var err error
	db, err = pgx.Connect(context.Background(), "postgres://user:pass@localhost:5432/shortener")
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	log.Println("Connected to PostgreSQL")
}

// SaveURL stores a URL mapping in the database
func SaveURL(shortCode, longURL string) {
	_, err := db.Exec(context.Background(), "INSERT INTO urls (short_code, long_url) VALUES ($1, $2)", shortCode, longURL)
	if err != nil {
		log.Println("DB Error:", err)
	}
}

// GetOriginalURL retrieves the original URL by its short code
func GetOriginalURL(shortCode string) string {
	var longURL string
	err := db.QueryRow(context.Background(), "SELECT long_url FROM urls WHERE short_code=$1", shortCode).Scan(&longURL)
	if err != nil {
		return ""
	}
	return longURL
}
