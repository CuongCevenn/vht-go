package categorymodule

import (
	categorycontroller "vht-go/modules/category/controller"
	categoryrepository "vht-go/modules/category/repository"
	categoryservice "vht-go/modules/category/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Dependencies Injection
func SetupCategoryModule(v1 * gin.RouterGroup, db *gorm.DB) {
	repo := categoryrepository.NewGORMCategoryRepository(db)
	service := categoryservice.NewCategoryService(repo)
	controller := categorycontroller.NewHTTPCategoryController(service)

	controller.SetupRoutes(v1)
}