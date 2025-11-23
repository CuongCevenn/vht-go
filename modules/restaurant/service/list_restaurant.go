package restaurantservice

import (
	"context"
	restaurantdomain "vht-go/modules/restaurant/domain"
	"vht-go/shared"
)

type ListRestaurantQuery struct {
	OwnerId *int          `form:"owner_id"`
	CityId  *int          `form:"city_id"`
	Status  *int          `form:"status"`
	Paging  shared.Paging `form:"paging"`
}

type ListRestaurantResult struct {
	Data   []restaurantdomain.Restaurant `json:"data"`
	Paging shared.Paging                 `json:"paging"`
}

func (s *RestaurantService) ListRestaurants(ctx context.Context, query *ListRestaurantQuery) (*ListRestaurantResult, error) {
	// Process paging defaults
	query.Paging.Process()

	// Check if filters are applied
	hasFilters := query.OwnerId != nil || query.CityId != nil || query.Status != nil

	var restaurants []restaurantdomain.Restaurant
	var total int64
	var err error

	if hasFilters {
		// Query with filters
		restaurants, err = s.repo.FindWithFilters(
			ctx,
			query.OwnerId,
			query.CityId,
			query.Status,
			(query.Paging.Page-1)*query.Paging.Limit,
			query.Paging.Limit,
		)
		if err != nil {
			return nil, err
		}

		total, err = s.repo.CountWithFilters(ctx, query.OwnerId, query.CityId, query.Status)
		if err != nil {
			return nil, err
		}
	} else {
		// Query all
		restaurants, err = s.repo.FindAll(
			ctx,
			(query.Paging.Page-1)*query.Paging.Limit,
			query.Paging.Limit,
		)
		if err != nil {
			return nil, err
		}

		total, err = s.repo.Count(ctx)
		if err != nil {
			return nil, err
		}
	}

	query.Paging.Total = total

	return &ListRestaurantResult{
		Data:   restaurants,
		Paging: query.Paging,
	}, nil
}

