package routes

import (
	"cutbray/test-gorm/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAppRoute(router *gin.RouterGroup) {
	router.GET("/hello", handlers.AppHello())

}
