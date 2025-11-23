package restaurantservice

import (
	"context"
	restaurantdomain "vht-go/modules/restaurant/domain"
)

type GetRestaurantQuery struct {
	Id int
}

type IGetRestaurantRepository interface {
	FindById(ctx context.Context, id int) (*restaurantdomain.Restaurant, error)
}

type GetRestaurantQueryHandler struct {
	repo IGetRestaurantRepository
}

func NewGetRestaurantQueryHandler(repo IGetRestaurantRepository) *GetRestaurantQueryHandler {
	return &GetRestaurantQueryHandler{repo: repo}
}

func (h *GetRestaurantQueryHandler) Handle(ctx context.Context, query *GetRestaurantQuery) (*restaurantdomain.Restaurant, error) {
	return h.repo.FindById(ctx, query.Id)
}

