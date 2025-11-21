package categorymodule

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var newCategory CreateCategoryDTO

		if err := c.ShouldBindJSON(&newCategory); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		newId, err := uuid.NewV7()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		createdCategory := Category{
			ID:          newId,
			Name:        newCategory.Name,
			Description: newCategory.Description,
			Status:      1,
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
		}
		
		if err := db.Create(&createdCategory).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"data": createdCategory.ID,
		})
	}

}