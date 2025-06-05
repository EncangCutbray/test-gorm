package routes

import (
	"github.com/gin-gonic/gin"

	docs "cutbray/test-gorm/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterSwaggerRoute(router *gin.Engine) {
	docs.SwaggerInfo.Title = "Cutbray Endpoint"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Description = "This is the API documentation for Cutbray. Use the endpoints to interact with the system."

	router.GET("/endpoint-doc/*any", ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/endpoint-doc/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)))
}
