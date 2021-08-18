package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s1ntaxe770r/opta-sanbox/db"
	"github.com/s1ntaxe770r/opta-sanbox/models"
)


// 	Insert
// @Summary Add a new value to the database
// @Description add a key value pair
// @ID insert
// @Accept  json
// @Produce  json
// @Param    	Request      body models.StorageRequest.Key  true  "Key"
// @Success 200 {object} models.StorageRequest
// @Failure 400 {string} string "Bad Request"
// @Router /set [post]
func Insert(c *gin.Context) {
	var request models.StorageRequest
	ctx := context.Background()
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"ERROR": err.Error(),
		})
		return
	}
	err = db.DB.Set(ctx, request.Key, request.Value, 0).Err()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}
	request.Status = "SUCCESS"
	c.JSON(http.StatusOK, gin.H{
		"SUCCESS": request,
	})

}
