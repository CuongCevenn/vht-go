package categorycontroller

import (
	"net/http"
	categorydomain "vht-go/modules/category/domain"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPCategoryController) GetListCategoriesAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging categorydomain.Paging

		if err := c.ShouldBindQuery(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}		

		paging.Process()

		var categories []categorydomain.Category
		categories, err := ctrl.svc.GetAllCategories(c.Request.Context(), &paging)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": categories,
			"paging": paging,
		})
	}
}