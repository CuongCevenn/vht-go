package ordermodule

import (
	ordercontroller "vht-go/modules/order/infras/controller"
	"vht-go/modules/order/infras/controller/orderrpcserver"
	orderrepository "vht-go/modules/order/infras/repository"
	"vht-go/modules/order/infras/repository/orderrpcclient"
	orderservice "vht-go/modules/order/service"
	"vht-go/shared"
	sharedcomponent "vht-go/shared/component"

	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
)

// Dependencies Injection
func SetupOrderModule(v1 *gin.RouterGroup, sctx sctx.ServiceContext) {
	appConfig := sctx.MustGet(sharedcomponent.AppConfigID).(sharedcomponent.IAppConfig)
	db := sctx.MustGet(shared.KeyGormComp).(sharedcomponent.IGormComp).DB()

	// 1. Initialize repository
	repo := orderrepository.NewGORMOrderRepository(db)

	// 2. Initialize RPC clients
	foodRPCClient := orderrpcclient.NewFoodRPCClient(appConfig.FoodServiceURI())
	userRPCClient := orderrpcclient.NewUserRPCClient(appConfig.UserServiceURI())

	// 3. Initialize handlers with repository and RPC clients
	createHandler := orderservice.NewCreateOrderResultCommandHandler(repo)
	getHandler := orderservice.NewGetOrderQueryHandler(repo, userRPCClient, foodRPCClient)
	listHandler := orderservice.NewListOrderQueryHandler(repo, userRPCClient, foodRPCClient)
	updateHandler := orderservice.NewUpdateOrderCommandHandler(repo)
	deleteHandler := orderservice.NewDeleteOrderCommandHandler(repo)

	// 4. Initialize controller with handlers
	controller := ordercontroller.NewHTTPOrderController(
		createHandler,
		getHandler,
		listHandler,
		updateHandler,
		deleteHandler,
	)

	// 5. Setup routes
	controller.SetupRoutes(v1)

	// 6. Setup RPC server
	rpcServer := orderrpcserver.NewOrderRPCServer(db)
	rpcServer.SetupRouter(v1)
}

