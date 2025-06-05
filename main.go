package main

import (
	"github.com/gin-gonic/gin"

	"cutbray/test-gorm/routes"
)

func main() {
	router := gin.Default()

	router.SetTrustedProxies(nil)

	routerV1 := router.Group("/api/v1")

	routes.RegisterAppRoute(routerV1)

	routes.RegisterSwaggerRoute(router)

	router.Run(":8080")
}
