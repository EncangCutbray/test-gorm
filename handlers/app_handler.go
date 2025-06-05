package handlers

import (
	"cutbray/test-gorm/request"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

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
func Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello From Server",
		})
	}
}

// @Summary Send message
// @Description Return Message From Server
// @Tags Example
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param content formData string true "Hello message content"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 422
// @Failure 500
// @Router /api/v1/message [post]
func Message() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req request.MessageHelloRequest

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := validate.Struct(req); err != nil {
			c.JSON(400, gin.H{"error": "validation failed", "detail": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": req})

	}
}
