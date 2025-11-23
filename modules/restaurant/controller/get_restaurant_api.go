package restaurantcontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPRestaurantController) GetRestaurantByIdAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id format"})
			return
		}

		restaurant, err := ctrl.svc.GetRestaurantById(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": restaurant})
	}
}

