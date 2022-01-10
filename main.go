package main

import (
	"github.com/hinchley2018/inventory-go/routes"

	"github.com/gin-gonic/gin"
	docs "github.com/hinchley2018/inventory-go/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	routes.CategoryRoutes(v1)
	routes.PartRoutes(v1)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
