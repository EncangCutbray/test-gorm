package routes

import (
	"cutbray/test-gorm/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAppRoute(router *gin.RouterGroup) {
	router.GET("/hello", handlers.Hello())
	router.GET("/hello-query", handlers.HelloQuery())
	router.POST("/hello-form-data", handlers.HelloFormData())
	router.POST("/hello-body", handlers.HelloBody())

}
