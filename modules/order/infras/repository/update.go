package orderrepository

import (
	"context"
	orderdomain "vht-go/modules/order/domain"
)

func (repo *GORMOrderRepository) Update(ctx context.Context, order *orderdomain.Order) error {
	return repo.db.WithContext(ctx).Where("id = ?", order.Id).Save(order).Error
}

