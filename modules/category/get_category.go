package categorymodule

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCategoryById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	for _, category := range categories {
		if category.ID == id {
			c.JSON(http.StatusOK, gin.H{"data": category})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "category not found"})
}	