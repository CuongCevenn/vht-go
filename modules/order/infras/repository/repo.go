package orderrepository

import "gorm.io/gorm"

type GORMOrderRepository struct {
	db *gorm.DB
}

func NewGORMOrderRepository(db *gorm.DB) *GORMOrderRepository {
	return &GORMOrderRepository{db: db}
}

