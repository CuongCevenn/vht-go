package categorymodule

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UpdateCategory(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		var catBody UpdateCategoryDTO

		if err := c.ShouldBindJSON(&catBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	

		if err := db.Where("id = ?", id).Updates(&catBody).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}	
		
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}