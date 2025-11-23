package categoryservice

import (
	"context"
	"errors"
	categorydomain "vht-go/modules/category/domain"
	"vht-go/shared"
)

type FilterStatusDTO struct {
	Status *int `json:"status,omitempty" form:"status,omitempty"`
}

func (dto *FilterStatusDTO) Validate() error {
	if dto.Status != nil {
		if *dto.Status < 0 || *dto.Status > 1 {
			return errors.New(categorydomain.ErrInvalidStatusFilter)
		}
	}
	return nil
}

type ListCategoryDTO struct {
	Paging *shared.Paging
	Filter *FilterStatusDTO
}

func (s *CategoryService) GetAllCategories(ctx context.Context, dto *ListCategoryDTO) (categories []categorydomain.Category, err error) {
	return s.repo.FindAll(ctx, dto)
}