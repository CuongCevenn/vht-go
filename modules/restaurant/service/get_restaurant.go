package restaurantservice

import (
	"context"
	restaurantdomain "vht-go/modules/restaurant/domain"
)

func (s *RestaurantService) GetRestaurantById(ctx context.Context, id int) (*restaurantdomain.Restaurant, error) {
	return s.repo.FindById(ctx, id)
}

