package categoryservice

import (
	"context"
	"time"
	categorydomain "vht-go/modules/category/domain"

	"github.com/google/uuid"
)

type UpdateCategoryDTO struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *int    `json:"status,omitempty"`
}

func (s *CategoryService) UpdateCategory(ctx context.Context, dto *UpdateCategoryDTO, id *uuid.UUID) error {

	oldCategory, err := s.repo.FindById(ctx, id)
	if err != nil {
		return err
	}

	category := &categorydomain.Category{
		Id: *id,
		Name: oldCategory.Name,
		Description: oldCategory.Description,
		Status: oldCategory.Status,
		CreatedAt: oldCategory.CreatedAt,
		UpdatedAt: time.Now(),
	}
	
	if dto.Name != nil {
		category.Name = *dto.Name
	}
	if dto.Description != nil {
		category.Description = *dto.Description
	}
	if dto.Status != nil {
		category.Status = *dto.Status
	}

	return s.repo.Update(ctx, category, id)
}