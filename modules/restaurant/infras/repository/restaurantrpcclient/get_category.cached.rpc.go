package restaurantrpcclient

import (
	"context"
	"fmt"
	restaurantdomain "vht-go/modules/restaurant/domain"
	sharedcomponent "vht-go/shared/component"

	"github.com/google/uuid"
)

// Proxy design pattern

type GetCategoryCachedRPCClient struct {
	originRepo *CategoryRPCClient
	redisComp  sharedcomponent.IRedisComp
}

func NewGetCategoryCachedRPCClient(originRepo *CategoryRPCClient, redisComp sharedcomponent.IRedisComp) *GetCategoryCachedRPCClient {
	return &GetCategoryCachedRPCClient{originRepo: originRepo, redisComp: redisComp}
}

func (c *GetCategoryCachedRPCClient) FindCategoryById(ctx context.Context, id uuid.UUID) (*restaurantdomain.RestaurantCategory, error) {
	var category restaurantdomain.RestaurantCategory

	key := fmt.Sprintf("category:%s", id.String())

	err := c.redisComp.Get(ctx, key, &category)

	if err != nil {
		category, err := c.originRepo.FindCategoryById(ctx, id)

		if err != nil {
			return nil, err
		}

		c.redisComp.Set(ctx, key, category, 0)
		return category, nil
	}

	return &category, nil
}