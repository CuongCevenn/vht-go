package orderdtos

import (
	"vht-go/shared"
)

type ListOrderDTO struct {
	UserId *string `json:"user_id,omitempty" form:"user_id"`
	FoodId *int    `json:"food_id,omitempty" form:"food_id"`
	Status *int    `json:"status,omitempty" form:"status"`
	Paging shared.Paging
}

