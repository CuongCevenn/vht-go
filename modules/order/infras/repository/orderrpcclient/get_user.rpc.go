package orderrpcclient

import (
	"context"
	"fmt"
	orderdomain "vht-go/modules/order/domain"

	"github.com/google/uuid"
	"resty.dev/v3"
)

func (c *UserRPCClient) FindUserById(ctx context.Context, id uuid.UUID) (*orderdomain.OrderUser, error) {
	fullURL := fmt.Sprintf("%s/get-user", c.userServiceURI)

	var dataRPC struct {
		Data orderdomain.OrderUser `json:"data"`
	}

	_, err := resty.New().R().
		SetBody(map[string]any{
			"id": id.String(),
		}).
		SetResult(&dataRPC).
		Post(fullURL)

	if err != nil {
		return nil, err
	}

	return &dataRPC.Data, nil
}

