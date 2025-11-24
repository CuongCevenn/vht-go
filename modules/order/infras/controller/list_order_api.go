package ordercontroller

import (
	"net/http"
	"strconv"

	orderdtos "vht-go/modules/order/dtos"
	orderservice "vht-go/modules/order/service"
	"vht-go/shared"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPOrderController) ListOrderAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto orderdtos.ListOrderDTO

		if err := c.ShouldBindQuery(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		// Parse pagination parameters
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

		paging := shared.Paging{
			Page:  page,
			Limit: limit,
		}

		query := &orderservice.ListOrderQuery{
			UserId: dto.UserId,
			FoodId: dto.FoodId,
			Status: dto.Status,
			Paging: &paging,
		}

		orders, err := ctrl.listHandler.Handle(c.Request.Context(), query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":   orders,
			"paging": paging,
		})
	}
}

