package categorymodule

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var newCategory CreateCategoryDTO

	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	lastId++
	createdCategory := Category{
		ID:          lastId,
		Name:        newCategory.Name,
		Description: newCategory.Description,
		Status:      1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	categories = append(categories, createdCategory)
	c.JSON(http.StatusCreated, createdCategory)
}