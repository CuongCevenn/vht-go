package restaurantcontroller

import (
	"net/http"
	restaurantservice "vht-go/modules/restaurant/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPRestaurantController) ListRestaurantsAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var query restaurantservice.ListRestaurantQuery

		if err := c.ShouldBindQuery(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		result, err := ctrl.svc.ListRestaurants(c.Request.Context(), &query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}

