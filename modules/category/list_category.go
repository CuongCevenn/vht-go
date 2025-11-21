package categorymodule

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListCategories(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging Paging

		if err := c.ShouldBindQuery(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}		

		paging.Process()

		var categories []Category
		if err := db.Table(Category{}.TableName()).Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).Order("id desc").Count(&paging.Total).Find(&categories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": categories,
			"paging": paging,
		})
	}
}