package cache

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func InitRedis() {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		log.Fatal("REDIS_ADDR is not set in environment variables")
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // No password by default
		DB:       0, // Default DB index
	})

	// Test the Redis connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("🚨 Redis connection failed: %v", err)
	}
	log.Println("✅ Connected to Redis")
}

// GetRedisClient returns the Redis client
func GetRedisClient() *redis.Client {
	return rdb
}
