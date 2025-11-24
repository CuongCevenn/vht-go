package orderdtos

import (
	"errors"
	orderdomain "vht-go/modules/order/domain"

	"github.com/google/uuid"
)

type CreateOrderDTO struct {
	UserId     uuid.UUID `json:"user_id" binding:"required"`
	FoodId     int       `json:"food_id" binding:"required"`
	Quantity   int       `json:"quantity" binding:"required"`
	TotalPrice float64   `json:"total_price" binding:"required"`
}

func (dto *CreateOrderDTO) Validate() error {
	if dto.UserId == uuid.Nil {
		return errors.New(orderdomain.ErrUserIdRequired)
	}

	if dto.FoodId <= 0 {
		return errors.New(orderdomain.ErrFoodIdRequired)
	}

	if dto.Quantity <= 0 {
		return errors.New(orderdomain.ErrQuantityInvalid)
	}

	if dto.TotalPrice <= 0 {
		return errors.New(orderdomain.ErrTotalPriceInvalid)
	}

	return nil
}

