package orderservice

import (
	"context"
)

type DeleteOrderCommand struct {
	Id int
}

type IDeleteOrderRepository interface {
	Delete(ctx context.Context, id int) error
}

type DeleteOrderCommandHandler struct {
	repo IDeleteOrderRepository
}

func NewDeleteOrderCommandHandler(repo IDeleteOrderRepository) *DeleteOrderCommandHandler {
	return &DeleteOrderCommandHandler{repo: repo}
}

func (h *DeleteOrderCommandHandler) Handle(ctx context.Context, cmd *DeleteOrderCommand) error {
	return h.repo.Delete(ctx, cmd.Id)
}

