package cache

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var rdb *redis.Client
var ctx = context.Background()

func InitRedis() {
	// Load .env file from config directory
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file in redis.go")
	}

	// Get the Redis URL from environment variables
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		log.Fatal("REDIS_URL is not set in .env")
	}

	// Connect to Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // Assuming no password for Redis
		DB:       0,  // Default DB
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Redis connection failed:", err)
	}
	log.Println("Connected to Redis")
}
