package ordercontroller

import (
	"net/http"

	orderdtos "vht-go/modules/order/dtos"
	orderservice "vht-go/modules/order/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPOrderController) CreateOrderAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto orderdtos.CreateOrderDTO

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		newId, err := ctrl.createHandler.Handle(c.Request.Context(), &orderservice.CreateOrderResultCommand{DTO: &dto})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"data": newId})
	}
}

