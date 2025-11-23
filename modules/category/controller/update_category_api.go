package categorycontroller

import (
	"net/http"
	categoryservice "vht-go/modules/category/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *HTTPCategoryController) UpdateCategoryAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		var dto categoryservice.UpdateCategoryDTO

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		cmd := categoryservice.UpdateCategoryCommand{
			Data: dto,
			Id:   id,
		}

		if err := ctrl.svc.UpdateCategory(c.Request.Context(), &cmd); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}