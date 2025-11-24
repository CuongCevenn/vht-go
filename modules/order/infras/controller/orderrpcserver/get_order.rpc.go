package orderrpcserver

import (
	orderdomain "vht-go/modules/order/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetOrderRPCRequest struct {
	Id int `json:"id" binding:"required"`
}

func (s *OrderRPCServer) GetOrderRPCAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetOrderRPCRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var order orderdomain.Order

		if err := s.db.First(&order, "id = ?", req.Id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(404, gin.H{"error": orderdomain.ErrOrderNotFound})
				return
			}
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": &order})
	}
}

