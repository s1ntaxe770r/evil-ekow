package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



// HealthCheck returns a 200 to indicate the service is alive
// @Summary use Anonymous field
// @Success 200
// @Failure 400
// @Failure 500
func HealthCheck(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}
