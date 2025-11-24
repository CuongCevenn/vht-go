package orderrpcserver

import (
	orderdomain "vht-go/modules/order/domain"

	"github.com/gin-gonic/gin"
)

type GetOrdersRPCRequest struct {
	Ids []int `json:"ids"`
}

func (s *OrderRPCServer) GetOrdersRPCAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetOrdersRPCRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var orders []orderdomain.Order

		if err := s.db.Where("id IN (?)", req.Ids).Find(&orders).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": orders})
	}

}

