package orderdtos

import (
	"errors"
	orderdomain "vht-go/modules/order/domain"
)

type UpdateOrderDTO struct {
	Quantity   *int     `json:"quantity,omitempty"`
	TotalPrice *float64 `json:"total_price,omitempty"`
	Status     *int     `json:"status,omitempty"`
}

func (dto *UpdateOrderDTO) Validate() error {
	if dto.Quantity != nil && *dto.Quantity <= 0 {
		return errors.New(orderdomain.ErrQuantityInvalid)
	}

	if dto.TotalPrice != nil && *dto.TotalPrice <= 0 {
		return errors.New(orderdomain.ErrTotalPriceInvalid)
	}

	return nil
}

