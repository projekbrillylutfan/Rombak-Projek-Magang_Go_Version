package app

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client


func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Alamat Redis server
		Password: "",               // Password jika ada
		DB:       0,                // Database default
	})
}

func SetRedisKey(key string, value string, expiration time.Duration) error {
	ctx := context.Background() // Context untuk Redis v9
	return RedisClient.Set(ctx, key, value, expiration).Err()
}

func GetRedisKey(key string) (string, error) {
	ctx := context.Background() // Context untuk Redis v9
	return RedisClient.Get(ctx, key).Result()
}