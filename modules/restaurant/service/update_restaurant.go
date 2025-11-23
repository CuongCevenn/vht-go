package restaurantservice

import (
	"context"
	"strings"
	"time"
	restaurantdomain "vht-go/modules/restaurant/domain"
)

type UpdateRestaurantDTO struct {
	Name             *string         `json:"name,omitempty"`
	Addr             *string         `json:"addr,omitempty"`
	CityId           *int            `json:"city_id,omitempty"`
	Lat              *float64        `json:"lat,omitempty"`
	Lng              *float64        `json:"lng,omitempty"`
	ShippingFeePerKm *float64        `json:"shipping_fee_per_km,omitempty"`
	Status           *int            `json:"status,omitempty"`
}

type UpdateRestaurantCommand struct {
	Id   int
	Data UpdateRestaurantDTO
}

func (s *RestaurantService) UpdateRestaurant(ctx context.Context, cmd *UpdateRestaurantCommand) error {
	// Fetch existing restaurant
	oldRestaurant, err := s.repo.FindById(ctx, cmd.Id)
	if err != nil {
		return err
	}

	// Build updated entity with existing values
	restaurant := &restaurantdomain.Restaurant{
		Id:               cmd.Id,
		OwnerId:          oldRestaurant.OwnerId,
		Name:             oldRestaurant.Name,
		Addr:             oldRestaurant.Addr,
		CityId:           oldRestaurant.CityId,
		Lat:              oldRestaurant.Lat,
		Lng:              oldRestaurant.Lng,
		ShippingFeePerKm: oldRestaurant.ShippingFeePerKm,
		Status:           oldRestaurant.Status,
		CreatedAt:        oldRestaurant.CreatedAt,
		UpdatedAt:        time.Now().UTC(),
	}

	// Apply changes with validation
	if cmd.Data.Name != nil {
		name := strings.TrimSpace(*cmd.Data.Name)
		if name != "" {
			restaurant.Name = name
		}
	}

	if cmd.Data.Addr != nil {
		addr := strings.TrimSpace(*cmd.Data.Addr)
		if addr != "" {
			restaurant.Addr = addr
		}
	}

	if cmd.Data.CityId != nil {
		restaurant.CityId = cmd.Data.CityId
	}

	if cmd.Data.Lat != nil {
		restaurant.Lat = cmd.Data.Lat
	}

	if cmd.Data.Lng != nil {
		restaurant.Lng = cmd.Data.Lng
	}
	
	if cmd.Data.ShippingFeePerKm != nil {
		if *cmd.Data.ShippingFeePerKm >= 0 {
			restaurant.ShippingFeePerKm = *cmd.Data.ShippingFeePerKm
		}
	}

	if cmd.Data.Status != nil {
		restaurant.Status = *cmd.Data.Status
	}

	// Persist changes
	return s.repo.Update(ctx, restaurant, cmd.Id)
}

