package categorymodule

import (
	categorycontroller "vht-go/modules/category/infras/controller"
	"vht-go/modules/category/infras/controller/categoryrpcserver"
	categoryrepository "vht-go/modules/category/infras/repository"
	categoryservice "vht-go/modules/category/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Dependencies Injection
func SetupCategoryModule(v1 * gin.RouterGroup, db *gorm.DB) {
	repo := categoryrepository.NewGORMCategoryRepository(db)
	// service := categoryservice.NewCategoryService(repo)
	createHandler := categoryservice.NewCreateCategoryResultCommandHandler(repo)
	getHandler := categoryservice.NewGetCategoryQueryHandler(repo)
	listHandler := categoryservice.NewListCategoryQueryHandler(repo)
	updateHandler := categoryservice.NewUpdateCategoryCommandHandler(repo)
	deleteHandler := categoryservice.NewDeleteCategoryCommandHandler(repo, repo)
	
	controller := categorycontroller.NewHTTPCategoryController(
		createHandler, 
		getHandler, 
		listHandler, 
		updateHandler,
		deleteHandler)

	controller.SetupRoutes(v1)

	rpcServer := categoryrpcserver.NewCategoryRPCServer(db)
	rpcServer.SetupRouter(v1)
}