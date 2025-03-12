package cache

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var rdb *redis.Client
var ctx = context.Background()

func InitRedis() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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


func CacheURL(shortURL, longURL string) {
	rdb.Set(ctx, shortURL, longURL, 24*time.Hour)
}

func GetCachedURL(shortURL string) (string, error) {
	val, err := rdb.Get(ctx, shortURL).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}
