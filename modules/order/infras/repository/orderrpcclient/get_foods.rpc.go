package orderrpcclient

import (
	"context"
	"fmt"
	orderdomain "vht-go/modules/order/domain"

	"resty.dev/v3"
)

func (c *FoodRPCClient) FindFoodsByIds(ctx context.Context, ids []int) ([]orderdomain.OrderFood, error) {
	fullURL := fmt.Sprintf("%s/get-foods", c.foodServiceURI)

	var dataRPC struct {
		Data []orderdomain.OrderFood `json:"data"`
	}

	_, err := resty.New().R().
		SetBody(map[string]any{
			"ids": ids,
		}).
		SetResult(&dataRPC).
		Post(fullURL)

	return dataRPC.Data, err
}

