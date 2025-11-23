package restaurantcontroller

import (
	"net/http"
	restaurantservice "vht-go/modules/restaurant/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPRestaurantController) CreateRestaurantAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto restaurantservice.CreateRestaurantDTO

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		newId, err := ctrl.svc.CreateNewRestaurant(c.Request.Context(), &dto)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": newId})
	}
}

