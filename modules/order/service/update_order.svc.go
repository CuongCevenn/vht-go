package orderservice

import (
	"context"
	"errors"
	"time"

	orderdomain "vht-go/modules/order/domain"
	orderdtos "vht-go/modules/order/dtos"
)

type UpdateOrderCommand struct {
	Id  int
	DTO *orderdtos.UpdateOrderDTO
}

type IUpdateOrderRepository interface {
	FindById(ctx context.Context, id int) (*orderdomain.Order, error)
	Update(ctx context.Context, order *orderdomain.Order) error
}

type UpdateOrderCommandHandler struct {
	repo IUpdateOrderRepository
}

func NewUpdateOrderCommandHandler(repo IUpdateOrderRepository) *UpdateOrderCommandHandler {
	return &UpdateOrderCommandHandler{repo: repo}
}

func (h *UpdateOrderCommandHandler) Handle(ctx context.Context, cmd *UpdateOrderCommand) error {
	if err := cmd.DTO.Validate(); err != nil {
		return err
	}

	// Find existing order
	order, err := h.repo.FindById(ctx, cmd.Id)
	if err != nil {
		return err
	}

	if order == nil {
		return errors.New(orderdomain.ErrOrderNotFound)
	}

	// Apply updates
	if cmd.DTO.Quantity != nil {
		order.Quantity = *cmd.DTO.Quantity
	}

	if cmd.DTO.TotalPrice != nil {
		order.TotalPrice = *cmd.DTO.TotalPrice
	}

	if cmd.DTO.Status != nil {
		order.Status = *cmd.DTO.Status
	}

	order.UpdatedAt = time.Now().UTC()

	return h.repo.Update(ctx, order)
}

