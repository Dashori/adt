package redisrepo

import (
	db "app/internal/db"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisRepo struct {
	TTL         time.Duration
	RedisClient db.RedisClient
}

func NewRedisRepo(redisClient db.RedisClient, ttl time.Duration) RedisRepository {
	return &RedisRepo{
		TTL:         ttl,
		RedisClient: redisClient,
	}
}

func (rr *RedisRepo) Set(key string, value string) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal by key %s, error is: %s", key, err)
	}

	if err := rr.RedisClient.Set(context.Background(), key, bytes, rr.TTL); err != nil {
		return fmt.Errorf("failed to set value to redis by key %s, error is: %s", key, err)
	}

	return nil
}

func (rr *RedisRepo) Get(key string) (string, error) {
	bytes, err := rr.RedisClient.Get(context.Background(), key)
	if err == redis.Nil {
		return "", err
	}

	if err != redis.Nil && err != nil {
		return "", fmt.Errorf("failed to get value from redis by key %s, error is: %s", key, err)
	}

	var value string

	if err := json.Unmarshal(bytes, &value); err != nil {
		return "", fmt.Errorf("failed to unmarshal by key %s, error is: %s", key, err)
	}

	return value, nil
}

func (rr *RedisRepo) Del(key string) error {
	err := rr.RedisClient.Del(context.Background(), key)

	if err != redis.Nil && err != nil {
		return fmt.Errorf("failed to delete pair by key %s, error is: %s", key, err)
	}

	return err
}
