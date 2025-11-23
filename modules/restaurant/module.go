package restaurantmodule

import (
	restaurantcontroller "vht-go/modules/restaurant/infras/controller"
	restaurantrepository "vht-go/modules/restaurant/infras/repository"
	restaurantservice "vht-go/modules/restaurant/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Dependencies Injection
func SetupRestaurantModule(v1 *gin.RouterGroup, db *gorm.DB) {
	// 1. Initialize repository
	repo := restaurantrepository.NewGORMRestaurantRepository(db)

	// 2. Initialize handlers with repository
	createHandler := restaurantservice.NewCreateRestaurantResultCommandHandler(repo)
	getHandler := restaurantservice.NewGetRestaurantQueryHandler(repo)
	listHandler := restaurantservice.NewListRestaurantQueryHandler(repo)
	updateHandler := restaurantservice.NewUpdateRestaurantCommandHandler(repo)
	deleteHandler := restaurantservice.NewDeleteRestaurantCommandHandler(repo, repo)

	// 3. Initialize controller with handlers
	controller := restaurantcontroller.NewHTTPRestaurantController(
		createHandler,
		getHandler,
		listHandler,
		updateHandler,
		deleteHandler,
	)

	// 4. Setup routes
	controller.SetupRoutes(v1)
}

