package restaurantmodule

import (
	restaurantcontroller "vht-go/modules/restaurant/controller"
	restaurantrepository "vht-go/modules/restaurant/repository"
	restaurantservice "vht-go/modules/restaurant/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRestaurantModule(v1 *gin.RouterGroup, db *gorm.DB) {
	// Initialize repository
	repo := restaurantrepository.NewGORMRestaurantRepository(db)

	// Initialize service with repository
	service := restaurantservice.NewRestaurantService(repo)

	// Initialize controller with service
	controller := restaurantcontroller.NewHTTPRestaurantController(service)

	// Setup routes
	controller.SetupRoutes(v1)
}

