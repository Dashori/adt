package middlewares

import (
	redisrepo "app/internal/repo"
	"github.com/gin-gonic/gin"
)

func RedisRepo(redisrepo redisrepo.RedisRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("redis_repo", redisrepo)
		c.Next()
	}
}
