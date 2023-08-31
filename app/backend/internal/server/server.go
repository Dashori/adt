package server

import (
	db "app/internal/db"
	redisrepo "app/internal/repo"
	controllers "app/internal/server/controllers"
	middlewares "app/internal/server/middlewares"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

var (
	redisRepo redisrepo.RedisRepository
)

func init() {
	ctx := context.Background()

	redisClient, err := db.NewRedisClient(ctx, &db.RedisOptions{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       os.Getenv("REDIS_DB"),
	})

	if err != nil {
		log.Fatalf("failed to create redis client, error is: %s", err)
	}

	redisRepo = redisrepo.NewRedisRepo(*redisClient, 1*time.Hour)
}

func SetupServer() *gin.Engine {

	router := gin.Default()

	router.Use(middlewares.RedisRepo(redisRepo))

	redisApi := router.Group("/")
	{
		redisApi.GET("/get_key", controllers.GetKey)
		redisApi.POST("/set_key", controllers.SetKey)
		redisApi.DELETE("/del_key", controllers.DelKey)
	}

	return router
}
