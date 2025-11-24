package orderrpcserver

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderRPCServer struct {
	db *gorm.DB
}

func NewOrderRPCServer(db *gorm.DB) *OrderRPCServer {
	return &OrderRPCServer{db: db}
}

func (s *OrderRPCServer) SetupRouter(v1 *gin.RouterGroup) {
	orderRPCGroup := v1.Group("/rpc/orders")
	orderRPCGroup.POST("/get-order", s.GetOrderRPCAPI())
	orderRPCGroup.POST("/get-orders", s.GetOrdersRPCAPI())
}

