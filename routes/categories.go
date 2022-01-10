package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryApiResponse struct {
	ID   int
	Name string
}
type CategoryApiRequest struct {
	Name string `binding:"required"`
}

// @BasePath /api/v1
// @Summary list categories
// @Accept json
// @Produce json
// @Success 200
// @Router /categories [get]
func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello",
	})
}

// @Summary create category
// @Accept json
// @Produce json
// @Success 200
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	var json CategoryApiRequest
	if err := c.ShouldBindJSON(&json); err != nil {

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello",
	})
}

// @Summary update category
// @Accept json
// @Produce json
// @Success 200
// @Router /categories/:id [put]
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var json CategoryApiRequest
	if err := c.ShouldBindJSON(&json); err != nil {

	}

	c.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}

func CategoryRoutes(rg *gin.RouterGroup) {
	categories := rg.Group("/categories")
	categories.GET("/", GetCategories)
	categories.POST("/", CreateCategory)
	categories.PUT("/:id", UpdateCategory)
}
