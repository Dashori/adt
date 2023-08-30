package db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"sync"
	"time"
	"crypto/tls"
	// "crypto/x509"
	// "os"
)

type RedisOptions struct {
	Host     string
	Port     string
	Password string
	DB       string
	TLSConfig tls.Config
}

type RedisClient struct {
	Client *redis.Client
}

var (
	redisClient *RedisClient
	lock        = &sync.Mutex{}
)

func NewRedisClient(ctx context.Context, options *RedisOptions) (*RedisClient, error) {
	lock.Lock()
	defer lock.Unlock()

	if redisClient == nil {
		client, err := getRedisClient(ctx, options)
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

func getRedisClient(ctx context.Context, options *RedisOptions) (*redis.Client, error) {
	// cert, err := tls.LoadX509KeyPair("../../../keys/app.crt", "../../../keys/app.key")
	// if err != nil {
	// 	return nil, fmt.Errorf("error is: %s", err)
	// }

	// // Load CA cert
	// caCert, err := os.ReadFile("../../../keys/ca.crt")
	// if err != nil {
	// 	return nil, fmt.Errorf("error is: %s", err)
	// }

	// caCertPool := x509.NewCertPool()
	// caCertPool.AppendCertsFromPEM(caCert)

	dbOpt, _ := strconv.Atoi(options.DB)

	opts := redis.Options{
		Addr:     fmt.Sprintf("%s:%s", options.Host, "6379"),
		Password: options.Password,
		DB:       dbOpt,
		// TLSConfig: &tls.Config{
		// 	MinVersion:   tls.VersionTLS12,
		// 	Certificates: []tls.Certificate{cert},
		// 	RootCAs:      caCertPool,
		// },
		// TLSConfig: &tls.Config{
		// 	InsecureSkipVerify: true,
		// },
	}
	client := redis.NewClient(&opts)

	if err := client.Ping(ctx).Err(); err != nil {
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

func (rc *RedisClient) Del(ctx context.Context, key string) error {
	return rc.Client.Del(ctx, key).Err()
}
