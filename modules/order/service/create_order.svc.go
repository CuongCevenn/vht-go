package orderservice

import (
	"context"
	"time"

	orderdomain "vht-go/modules/order/domain"
	orderdtos "vht-go/modules/order/dtos"
)

type CreateOrderResultCommand struct {
	DTO *orderdtos.CreateOrderDTO
}

type ICreateOrderRepository interface {
	Insert(ctx context.Context, order *orderdomain.Order) error
}

type CreateOrderResultCommandHandler struct {
	repo ICreateOrderRepository
}

func NewCreateOrderResultCommandHandler(repo ICreateOrderRepository) *CreateOrderResultCommandHandler {
	return &CreateOrderResultCommandHandler{repo: repo}
}

func (h *CreateOrderResultCommandHandler) Handle(ctx context.Context, cmd *CreateOrderResultCommand) (*int, error) {
	if err := cmd.DTO.Validate(); err != nil {
		return nil, err
	}

	order := orderdomain.Order{
		UserId:     cmd.DTO.UserId,
		FoodId:     cmd.DTO.FoodId,
		Quantity:   cmd.DTO.Quantity,
		TotalPrice: cmd.DTO.TotalPrice,
		Status:     1,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
	}

	if err := h.repo.Insert(ctx, &order); err != nil {
		return nil, err
	}

	return &order.Id, nil
}

