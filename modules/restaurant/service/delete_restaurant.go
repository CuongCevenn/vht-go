package restaurantservice

import (
	"context"
)

func (s *RestaurantService) DeleteRestaurant(ctx context.Context, id int) error {
	// Check if restaurant exists
	_, err := s.repo.FindById(ctx, id)
	if err != nil {
		return err
	}

	// Delete restaurant
	return s.repo.Delete(ctx, id)
}

