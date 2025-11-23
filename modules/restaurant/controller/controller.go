package restaurantcontroller

import (
	restaurantservice "vht-go/modules/restaurant/service"

	"github.com/gin-gonic/gin"
)

type HTTPRestaurantController struct {
	svc restaurantservice.IRestaurantService
}

func NewHTTPRestaurantController(svc restaurantservice.IRestaurantService) *HTTPRestaurantController {
	return &HTTPRestaurantController{svc: svc}
}

func (ctrl *HTTPRestaurantController) SetupRoutes(v1 *gin.RouterGroup) {
	restaurantGroup := v1.Group("/restaurants")
	restaurantGroup.POST("", ctrl.CreateRestaurantAPI())
	restaurantGroup.GET("/:id", ctrl.GetRestaurantByIdAPI())
	restaurantGroup.PATCH("/:id", ctrl.UpdateRestaurantAPI())
	restaurantGroup.GET("", ctrl.ListRestaurantsAPI())
	restaurantGroup.DELETE("/:id", ctrl.DeleteRestaurantAPI())
}

