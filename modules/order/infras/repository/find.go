package orderrepository

import (
	"context"
	"errors"

	orderdomain "vht-go/modules/order/domain"
	"gorm.io/gorm"
)

func (repo *GORMOrderRepository) FindById(ctx context.Context, id int) (*orderdomain.Order, error) {
	var order orderdomain.Order
	if err := repo.db.WithContext(ctx).Where("id = ?", id).First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(orderdomain.ErrOrderNotFound)
		}
		return nil, err
	}
	return &order, nil
}

func (repo *GORMOrderRepository) FindAll(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]orderdomain.Order, error) {
	var orders []orderdomain.Order
	query := repo.db.WithContext(ctx)

	// Apply filters
	if userId, ok := filter["user_id"]; ok && userId != nil {
		query = query.Where("user_id = ?", userId)
	}

	if foodId, ok := filter["food_id"]; ok && foodId != nil {
		query = query.Where("food_id = ?", foodId)
	}

	if status, ok := filter["status"]; ok && status != nil {
		query = query.Where("status = ?", status)
	}

	// Apply pagination
	query = query.Limit(limit).Offset(offset)

	if err := query.Order("created_at DESC").Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (repo *GORMOrderRepository) Count(ctx context.Context, filter map[string]interface{}) (int64, error) {
	var count int64
	query := repo.db.WithContext(ctx).Model(&orderdomain.Order{})

	// Apply filters
	if userId, ok := filter["user_id"]; ok && userId != nil {
		query = query.Where("user_id = ?", userId)
	}

	if foodId, ok := filter["food_id"]; ok && foodId != nil {
		query = query.Where("food_id = ?", foodId)
	}

	if status, ok := filter["status"]; ok && status != nil {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

