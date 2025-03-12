package cache

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
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
