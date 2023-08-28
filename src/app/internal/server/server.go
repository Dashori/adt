package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAllDoctors(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"test": "pass"})
}

func SetupServer() *gin.Engine {

	router := gin.Default()

	api := router.Group("/api/v2")
	{
		api.GET("/test", getAllDoctors)
	}

	return router
}
