package orderrepository

import (
	"context"
	orderdomain "vht-go/modules/order/domain"
)

func (repo *GORMOrderRepository) Insert(ctx context.Context, order *orderdomain.Order) error {
	return repo.db.WithContext(ctx).Create(order).Error
}

