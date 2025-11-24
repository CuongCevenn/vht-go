package ordercontroller

import (
	"net/http"
	"strconv"

	orderservice "vht-go/modules/order/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPOrderController) GetOrderByIdAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid order id"})
			return
		}

		order, err := ctrl.getHandler.Handle(c.Request.Context(), &orderservice.GetOrderQuery{Id: id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": order})
	}
}

