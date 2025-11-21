package categorymodule

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var catBody UpdateCategoryDTO

	if err := c.ShouldBindJSON(&catBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	for i, category := range categories {
		if category.ID == id {
			if catBody.Name != nil {
				category.Name = *catBody.Name
			}
			if catBody.Description != nil {
				category.Description = *catBody.Description
			}
			if catBody.Status != nil {
				category.Status = *catBody.Status
			}
			category.UpdatedAt = time.Now()
			categories[i] = category
			c.JSON(http.StatusOK, category)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "category not found"})

}