package categorymodule

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	for i, category := range categories {
		if category.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"data": true})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "category not found"})
}