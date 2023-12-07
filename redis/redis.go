package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type IRedis interface {
	Set(key string, value string, ttl time.Duration) error
	Get(key string) (string, error)
}

type Redis struct {
	redis *redis.Client
}

func NewRedis() IRedis {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "root",
		DB:       0,
	})

	return &Redis{
		redis: client,
	}
}

func (r *Redis) Set(key string, value string, ttl time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.redis.Set(ctx, key, value, ttl).Err()
}

func (r *Redis) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.redis.Get(ctx, key).Result()
}
