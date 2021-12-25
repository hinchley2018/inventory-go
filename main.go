package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	docs "github.com/hinchley2018/inventory-go/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1
// @Summary list parts
// @Accept json
// @Produce json
// @Success 200
// @Router /categories [get]
func GetPartCategories(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello",
	})
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		categoriesGroup := v1.Group("/categories")
		{
			categoriesGroup.GET("/", GetPartCategories)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
