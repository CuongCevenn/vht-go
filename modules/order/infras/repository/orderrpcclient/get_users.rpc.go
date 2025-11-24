package orderrpcclient

import (
	"context"
	"fmt"
	orderdomain "vht-go/modules/order/domain"

	"github.com/google/uuid"
	"resty.dev/v3"
)

func (c *UserRPCClient) FindUsersByIds(ctx context.Context, ids []uuid.UUID) ([]orderdomain.OrderUser, error) {
	fullURL := fmt.Sprintf("%s/get-users", c.userServiceURI)

	var dataRPC struct {
		Data []orderdomain.OrderUser `json:"data"`
	}

	// Convert UUIDs to strings
	idStrings := make([]string, len(ids))
	for i, id := range ids {
		idStrings[i] = id.String()
	}

	_, err := resty.New().R().
		SetBody(map[string]any{
			"ids": idStrings,
		}).
		SetResult(&dataRPC).
		Post(fullURL)

	return dataRPC.Data, err
}

