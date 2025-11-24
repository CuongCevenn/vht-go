package orderservice

import (
	"context"

	orderdomain "vht-go/modules/order/domain"
	"vht-go/shared"

	"github.com/google/uuid"
)

type ListOrderQuery struct {
	UserId *string
	FoodId *int
	Status *int
	Paging *shared.Paging
}

type IListOrderRepository interface {
	FindAll(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]orderdomain.Order, error)
	Count(ctx context.Context, filter map[string]interface{}) (int64, error)
}

type IGetOrderUsers interface {
	FindUsersByIds(ctx context.Context, ids []uuid.UUID) ([]orderdomain.OrderUser, error)
}

type IGetOrderFoods interface {
	FindFoodsByIds(ctx context.Context, ids []int) ([]orderdomain.OrderFood, error)
}

type ListOrderQueryHandler struct {
	repo     IListOrderRepository
	userRepo IGetOrderUsers
	foodRepo IGetOrderFoods
}

func NewListOrderQueryHandler(repo IListOrderRepository, userRepo IGetOrderUsers, foodRepo IGetOrderFoods) *ListOrderQueryHandler {
	return &ListOrderQueryHandler{repo: repo, userRepo: userRepo, foodRepo: foodRepo}
}

func (h *ListOrderQueryHandler) Handle(ctx context.Context, query *ListOrderQuery) ([]orderdomain.Order, error) {
	query.Paging.Process()

	filter := make(map[string]interface{})
	if query.UserId != nil {
		filter["user_id"] = *query.UserId
	}
	if query.FoodId != nil {
		filter["food_id"] = *query.FoodId
	}
	if query.Status != nil {
		filter["status"] = *query.Status
	}

	// Get total count
	total, err := h.repo.Count(ctx, filter)
	if err != nil {
		return nil, err
	}
	query.Paging.Total = total

	// Get orders
	offset := (query.Paging.Page - 1) * query.Paging.Limit
	orders, err := h.repo.FindAll(ctx, filter, query.Paging.Limit, offset)
	if err != nil {
		return nil, err
	}

	// Collect user IDs and food IDs
	userIds := make([]uuid.UUID, 0)
	foodIds := make([]int, 0)
	userIdMap := make(map[uuid.UUID]bool)
	foodIdMap := make(map[int]bool)

	for _, order := range orders {
		if !userIdMap[order.UserId] {
			userIds = append(userIds, order.UserId)
			userIdMap[order.UserId] = true
		}
		if !foodIdMap[order.FoodId] {
			foodIds = append(foodIds, order.FoodId)
			foodIdMap[order.FoodId] = true
		}
	}

	// Fetch users
	var userMap map[uuid.UUID]orderdomain.OrderUser
	if len(userIds) > 0 {
		users, err := h.userRepo.FindUsersByIds(ctx, userIds)
		if err == nil {
			userMap = make(map[uuid.UUID]orderdomain.OrderUser)
			for _, user := range users {
				userMap[user.Id] = user
			}
		}
	}

	// Fetch foods
	var foodMap map[int]orderdomain.OrderFood
	if len(foodIds) > 0 {
		foods, err := h.foodRepo.FindFoodsByIds(ctx, foodIds)
		if err == nil {
			foodMap = make(map[int]orderdomain.OrderFood)
			for _, food := range foods {
				foodMap[food.Id] = food
			}
		}
	}

	// Populate orders with user and food data
	for i := range orders {
		if user, exists := userMap[orders[i].UserId]; exists {
			orders[i].User = &user
		}
		if food, exists := foodMap[orders[i].FoodId]; exists {
			orders[i].Food = &food
		}
	}

	return orders, nil
}

