package ordercontroller

import (
	"net/http"
	"strconv"

	orderdtos "vht-go/modules/order/dtos"
	orderservice "vht-go/modules/order/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPOrderController) UpdateOrderAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid order id"})
			return
		}

		var dto orderdtos.UpdateOrderDTO

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		err = ctrl.updateHandler.Handle(c.Request.Context(), &orderservice.UpdateOrderCommand{
			Id:  id,
			DTO: &dto,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

