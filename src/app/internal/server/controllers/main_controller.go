package controllers

import (
	redisrepo "app/internal/repo"
	models "app/internal/server/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
)

func GetKey(c *gin.Context) {
	redisRepo, ok := c.MustGet("redis_repo").(redisrepo.RedisRepository)
	if ok != true {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
			Message: "Failed to get redis_repo. Please, try again later"})
		return
	}

	var key string

	key, ok = c.GetQuery("key")
	if ok != true {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Message: "Failed to parse request body"})
		return
	}

	result, err := redisRepo.Get(key)
	if err == redis.Nil {
		c.AbortWithStatusJSON(http.StatusNotFound, models.Response{
			Message: "No value with this key in redis"})
		return
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ResponseError{
			Message: "Failed to get key from redis",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Success get value from redis",
		Key:     key,
		Value:   result,
	})
}

func SetKey(c *gin.Context) {

	redisRepo, ok := c.MustGet("redis_repo").(redisrepo.RedisRepository)
	if ok != true {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
			Message: "Failed to get redis_repo. Please, try again later"})
		return
	}

	var params models.Pair

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ResponseError{
			Message: "Failed to parse request body",
			Error:   err.Error(),
		})
		return
	}

	for key, value := range params {
		err = redisRepo.Set(key, value)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, models.ResponseError{
				Message: "Failed to set new key-value into redis",
				Error:   err.Error(),
			})
			return
		}

		// in case we have one and only one key-value pair and it's not empty
		// so we can return 200 here with key-value in json
		c.JSON(http.StatusOK, models.ResponseSuccess{
			Message: "Success add key-value into redis",
			Key:     key,
			Value:   value,
		})
	}
}

func DelKey(c *gin.Context) {
	redisRepo, ok := c.MustGet("redis_repo").(redisrepo.RedisRepository)
	if ok != true {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
			Message: "Failed to get redis_repo. Please, try again later"})
		return
	}

	var params models.Pair

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ResponseError{
			Message: "Failed to parse request body",
			Error:   err.Error(),
		})
		return
	}

	for _, value := range params {
		err = redisRepo.Del(value)
		
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, models.ResponseError{
				Message: "Failed to delete key-value from redis.",
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, models.ResponseSuccess{
			Message: "Success delete key-value from redis",
			Key:     value,
		})
	}
}
