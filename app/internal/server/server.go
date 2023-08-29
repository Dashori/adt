package server

import (
	redisrepo "app/internal/repo"
	controllers "app/internal/server/controllers"
	middlewares "app/internal/server/middlewares"
	flags "app/src/flags"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var (
	redisRepo redisrepo.RedisRepository
)

func SetupServer(options *flags.RedisFlags) *gin.Engine {
	redisClient, err := flags.NewRedisClient(options)

	if err != nil {
		log.Fatalf("failed to create redis client, error is: %s", err)
	}

	redisRepo = redisrepo.NewRedisRepo(*redisClient, 1*time.Hour)

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
