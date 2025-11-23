package categoryservice

import (
	"context"

	"github.com/google/uuid"
)

func (svc *CategoryService) DeleteCategory(ctx context.Context, id *uuid.UUID) error {
	if err := svc.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}