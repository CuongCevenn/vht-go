package categorymodule

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetCategoryById(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
	
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		var category Category
		if err := db.First(&category, "id = ?", id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"message": "category not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})	
			return
		}

		c.JSON(http.StatusOK, category)
	}
}	