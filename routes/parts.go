package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PartApiResponse struct {
	ID           int
	Name         string
	SerialNumber string
}
type PartApiRequest struct {
	Name         string `binding:"required"`
	SerialNumber string `binding:"required"`
}

// @BasePath /api/v1
// @Summary list parts
// @Accept json
// @Produce json
// @Success 200
// @Router /parts [get]
func GetParts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/parts",
	})
}

// @Summary create part
// @Accept json
// @Produce json
// @Success 200
// @Router /parts [post]
func CreatePart(c *gin.Context) {
	var json CategoryApiRequest
	if err := c.ShouldBindJSON(&json); err != nil {

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "POST /parts",
	})
}

// @Summary update category
// @Accept json
// @Produce json
// @Success 200
// @Router /categories/:id [put]
func UpdatePart(c *gin.Context) {
	id := c.Param("id")
	var json CategoryApiRequest
	if err := c.ShouldBindJSON(&json); err != nil {

	}

	c.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}

func PartRoutes(rg *gin.RouterGroup) {
	parts := rg.Group("/parts")
	parts.GET("/", GetParts)
	parts.POST("/", CreatePart)
	parts.PUT("/:id", UpdatePart)
}
