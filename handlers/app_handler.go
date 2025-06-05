package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Ping endpoint
// @Description Return Hello From Server
// @Tags Example
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 422
// @Failure 500
// @Router /api/v1/hello [get]
func AppHello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello From Server",
		})
	}
}
