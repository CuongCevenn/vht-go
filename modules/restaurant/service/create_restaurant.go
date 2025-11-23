package restaurantservice

import (
	"context"
	"strings"
	"time"
	restaurantdomain "vht-go/modules/restaurant/domain"
)

type CreateRestaurantDTO struct {
	OwnerId          int            `json:"owner_id" binding:"required"`
	Name             string         `json:"name" binding:"required"`
	Addr             string         `json:"addr" binding:"required"`
	CityId           *int           `json:"city_id,omitempty"`
	Lat              *float64       `json:"lat,omitempty"`
	Lng              *float64       `json:"lng,omitempty"`
	ShippingFeePerKm float64        `json:"shipping_fee_per_km"`
	Status           int            `json:"status"`
}

func (s *RestaurantService) CreateNewRestaurant(ctx context.Context, dto *CreateRestaurantDTO) (int, error) {
	// Validate and sanitize input
	name := strings.TrimSpace(dto.Name)
	if name == "" {
		return 0, nil
	}

	addr := strings.TrimSpace(dto.Addr)
	if addr == "" {
		return 0, nil
	}

	// Set default status if not provided
	status := dto.Status
	if status == 0 {
		status = 1
	}

	// Set default shipping fee if not provided
	shippingFee := dto.ShippingFeePerKm
	if shippingFee < 0 {
		shippingFee = 0
	}

	restaurant := &restaurantdomain.Restaurant{
		OwnerId:          dto.OwnerId,
		Name:             name,
		Addr:             addr,
		CityId:           dto.CityId,
		Lat:              dto.Lat,
		Lng:              dto.Lng,
		ShippingFeePerKm: shippingFee,
		Status:           status,
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
	}

	if err := s.repo.Insert(ctx, restaurant); err != nil {
		return 0, err
	}

	return restaurant.Id, nil
}

