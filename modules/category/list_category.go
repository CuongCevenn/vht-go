package categorymodule

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCategories(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": categories})
}