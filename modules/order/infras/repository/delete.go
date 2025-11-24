package orderrepository

import (
	"context"
	orderdomain "vht-go/modules/order/domain"
)

func (repo *GORMOrderRepository) Delete(ctx context.Context, id int) error {
	return repo.db.WithContext(ctx).Where("id = ?", id).Delete(&orderdomain.Order{}).Error
}

