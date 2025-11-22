package categorycontroller

import (
	"net/http"
	categorydomain "vht-go/modules/category/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *HTTPCategoryController) GetCategoryByIdAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
	
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		var category *categorydomain.Category
		category, err = ctrl.svc.GetCategoryById(c.Request.Context(), &id)

		if err != nil {
			if err.Error() == categorydomain.ErrCategoryNotFound {
				c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, category)
	}
}