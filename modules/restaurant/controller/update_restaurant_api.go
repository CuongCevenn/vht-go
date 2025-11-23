package restaurantcontroller

import (
	"net/http"
	"strconv"
	restaurantservice "vht-go/modules/restaurant/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPRestaurantController) UpdateRestaurantAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id format"})
			return
		}

		var dto restaurantservice.UpdateRestaurantDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		cmd := restaurantservice.UpdateRestaurantCommand{Id: id, Data: dto}
		if err := ctrl.svc.UpdateRestaurant(c.Request.Context(), &cmd); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

