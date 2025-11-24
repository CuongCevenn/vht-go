package orderrpcclient

import (
	"context"
	"fmt"
	orderdomain "vht-go/modules/order/domain"

	"resty.dev/v3"
)

func (c *FoodRPCClient) FindFoodById(ctx context.Context, id int) (*orderdomain.OrderFood, error) {
	fullURL := fmt.Sprintf("%s/get-food", c.foodServiceURI)

	var dataRPC struct {
		Data orderdomain.OrderFood `json:"data"`
	}

	_, err := resty.New().R().
		SetBody(map[string]any{
			"id": id,
		}).
		SetResult(&dataRPC).
		Post(fullURL)

	if err != nil {
		return nil, err
	}

	return &dataRPC.Data, nil
}

