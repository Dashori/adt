package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAllDoctors(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"test": "pass"})
}

func getKey(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"get key": "ok"})
}

func setKey(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"set key": "ok"})
}

func delKey(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"del key": "ok"})
}

func SetupServer() *gin.Engine {

	router := gin.Default()

	api := router.Group("/api/v2")
	{
		api.GET("/test", getAllDoctors)
	}

	redisApi := router.Group("/")
	{
		redisApi.GET("/get_key", getKey)
		redisApi.GET("/set_key", setKey)
		redisApi.GET("/del_key", delKey)
	}

	return router
}
