package ordercontroller

import (
	orderdomain "vht-go/modules/order/domain"
	orderservice "vht-go/modules/order/service"
	"vht-go/shared"

	"github.com/gin-gonic/gin"
)

type HTTPOrderController struct {
	createHandler shared.ICommandResultHandler[*orderservice.CreateOrderResultCommand, *int]
	getHandler    shared.IQueryHandler[*orderservice.GetOrderQuery, *orderdomain.Order]
	listHandler   shared.IQueryHandler[*orderservice.ListOrderQuery, []orderdomain.Order]
	updateHandler shared.ICommandHandler[*orderservice.UpdateOrderCommand]
	deleteHandler shared.ICommandHandler[*orderservice.DeleteOrderCommand]
}

func NewHTTPOrderController(
	createHandler shared.ICommandResultHandler[*orderservice.CreateOrderResultCommand, *int],
	getHandler shared.IQueryHandler[*orderservice.GetOrderQuery, *orderdomain.Order],
	listHandler shared.IQueryHandler[*orderservice.ListOrderQuery, []orderdomain.Order],
	updateHandler shared.ICommandHandler[*orderservice.UpdateOrderCommand],
	deleteHandler shared.ICommandHandler[*orderservice.DeleteOrderCommand]) *HTTPOrderController {
	return &HTTPOrderController{
		createHandler: createHandler,
		getHandler:    getHandler,
		listHandler:   listHandler,
		updateHandler: updateHandler,
		deleteHandler: deleteHandler,
	}
}

func (ctrl *HTTPOrderController) SetupRoutes(v1 *gin.RouterGroup) {
	orderGroup := v1.Group("/orders")
	{
		orderGroup.POST("", ctrl.CreateOrderAPI())
		orderGroup.GET("/:id", ctrl.GetOrderByIdAPI())
		orderGroup.GET("", ctrl.ListOrderAPI())
		orderGroup.PATCH("/:id", ctrl.UpdateOrderAPI())
		orderGroup.DELETE("/:id", ctrl.DeleteOrderAPI())
	}
}

