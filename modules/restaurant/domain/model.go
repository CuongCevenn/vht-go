package restaurantdomain

import (
	"time"

	"github.com/google/uuid"
)

type Restaurant struct {
	Id                int            `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	CategoryId		*uuid.UUID      `json:"category_id" gorm:"column:category_id"`
	OwnerId           int            `json:"owner_id" gorm:"column:owner_id;not null"`
	Name              string         `json:"name" gorm:"column:name;not null"`
	Addr              string         `json:"addr" gorm:"column:addr;not null"`
	CityId            *int           `json:"city_id" gorm:"column:city_id"`
	Lat               *float64       `json:"lat" gorm:"column:lat"`
	Lng               *float64       `json:"lng" gorm:"column:lng"`
	ShippingFeePerKm  float64        `json:"shipping_fee_per_km" gorm:"column:shipping_fee_per_km;default:0"`
	Status            int            `json:"status" gorm:"column:status;not null;default:1"`
	CreatedAt         time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt         time.Time      `json:"updated_at" gorm:"column:updated_at"`
	Category  	 *RestaurantCategory       `json:"category" gorm:"-"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCategory struct {
	Id          uuid.UUID `json:"id" gorm:"column:id"`
	Name        string    `json:"name" gorm:"column:name;"`
}

func (RestaurantCategory) TableName() string {
	return "categories"
}
