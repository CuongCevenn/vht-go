package orderservice

import (
	"context"

	orderdomain "vht-go/modules/order/domain"

	"github.com/google/uuid"
)

type GetOrderQuery struct {
	Id int
}

type IGetOrderRepository interface {
	FindById(ctx context.Context, id int) (*orderdomain.Order, error)
}

type IGetOrderUser interface {
	FindUserById(ctx context.Context, id uuid.UUID) (*orderdomain.OrderUser, error)
}

type IGetOrderFood interface {
	FindFoodById(ctx context.Context, id int) (*orderdomain.OrderFood, error)
}

type GetOrderQueryHandler struct {
	repo     IGetOrderRepository
	userRepo IGetOrderUser
	foodRepo IGetOrderFood
}

func NewGetOrderQueryHandler(repo IGetOrderRepository, userRepo IGetOrderUser, foodRepo IGetOrderFood) *GetOrderQueryHandler {
	return &GetOrderQueryHandler{repo: repo, userRepo: userRepo, foodRepo: foodRepo}
}

func (h *GetOrderQueryHandler) Handle(ctx context.Context, query *GetOrderQuery) (*orderdomain.Order, error) {
	order, err := h.repo.FindById(ctx, query.Id)

	if err != nil {
		return nil, err
	}

	// Populate user if exists
	if order != nil {
		user, err := h.userRepo.FindUserById(ctx, order.UserId)
		if err == nil {
			order.User = user
		}
	}

	// Populate food
	if order != nil {
		food, err := h.foodRepo.FindFoodById(ctx, order.FoodId)
		if err == nil {
			order.Food = food
		}
	}

	return order, nil
}

