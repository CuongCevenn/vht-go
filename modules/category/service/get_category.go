package categoryservice

import (
	"context"
	categorydomain "vht-go/modules/category/domain"

	"github.com/google/uuid"
)

func (s *CategoryService) GetCategoryById(ctx context.Context, id *uuid.UUID) (*categorydomain.Category, error) {
	return s.repo.FindById(ctx, id)
}