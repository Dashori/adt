package flags

import (
	"context"
	"sync"
	"github.com/redis/go-redis/v9"
	"fmt"
	"time"
)


type RedisFlags struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}

type RedisClient struct {
	Client *redis.Client
}

var (
	redisClient *RedisClient
	lock        = &sync.Mutex{}
)

func NewRedisClient(options *RedisFlags) (*RedisClient, error) {
	lock.Lock()
	defer lock.Unlock()

	if redisClient == nil {
		client, err := getRedisClient(options)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize redis client, error is: %s", err)
		}
		redisClient = &RedisClient{
			Client: client,
		}
		return redisClient, nil
	}

	return redisClient, nil
}

func getRedisClient(options *RedisFlags) (*redis.Client, error) {
	opts := redis.Options{
		Addr:     fmt.Sprintf("%s:%s", options.Host, options.Port),
		Password: options.Password,
	}
	client := redis.NewClient(&opts)

	if err := client.Ping(context.TODO()).Err(); err != nil {
		return nil, fmt.Errorf("redis client with ttl failed to ping address %s, error is: %s",
			opts.Addr, err)
	}

	return client, nil
}

func (rc *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return rc.Client.Set(ctx, key, value, expiration).Err()
}

func (rc *RedisClient) Get(ctx context.Context, key string) ([]byte, error) {
	return rc.Client.Get(ctx, key).Bytes()
}
