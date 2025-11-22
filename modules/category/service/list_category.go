package categoryservice

import (
	"context"
	categorydomain "vht-go/modules/category/domain"
)

func (s *CategoryService) GetAllCategories(ctx context.Context, paging *categorydomain.Paging) (categories []categorydomain.Category, err error) {
	return s.repo.FindAll(ctx, paging)
}