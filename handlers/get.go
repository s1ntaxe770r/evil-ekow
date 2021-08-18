package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/s1ntaxe770r/opta-sanbox/db"
	"github.com/s1ntaxe770r/opta-sanbox/models"
)


// Get returns the corresponding value of the key passed in the request
// @Summary Retrieve a value
// @Description returns the corresponding value for the key
// @ID Get
// @Accept  json
// @Produce  json
// @Param   Request    body models.StorageRequest true  "Some ID"
// @Success 200 {object} models.GetRequest
// @Failure 400 {string} string "Bad Request"
// @Router /get [post]
func Get(c *gin.Context) {
	request := new(models.GetRequest)
	ctx := context.Background()
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"ERROR": err.Error(),
		})
	}
	value, err := db.DB.Get(ctx, request.Key).Result()
	switch {
	case err == redis.Nil:
		c.JSON(404, gin.H{"ERROR": "key does not exist"})
	case err != nil:
		c.JSON(404, gin.H{"ERROR": err.Error()})
	}
	request.Value = value
	c.JSON(200, gin.H{
		"SUCCESS": request,
	})

}
