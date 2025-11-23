package categoryservice

import (
	"context"
	"errors"
	"strings"
	"time"
	categorydomain "vht-go/modules/category/domain"

	"github.com/google/uuid"
)

type UpdateCategoryDTO struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *int    `json:"status,omitempty"`
}

func (UpdateCategoryDTO) TableName() string {
	return categorydomain.Category{}.TableName()
}

type UpdateCategoryCommand struct {
	Id uuid.UUID `json:"id"`
	Data UpdateCategoryDTO `json:"data"`
}


func (s *CategoryService) UpdateCategory(ctx context.Context, cmd *UpdateCategoryCommand) error {

	oldCategory, err := s.repo.FindById(ctx, &cmd.Id)
	if err != nil {
		return err
	}

	category := &categorydomain.Category{
		Id: cmd.Id,
		Name: oldCategory.Name,
		Description: oldCategory.Description,
		Status: oldCategory.Status,
		CreatedAt: oldCategory.CreatedAt,
		UpdatedAt: time.Now(),
	}
	
	if cmd.Data.Name != nil {
		category.Name = strings.TrimSpace(*cmd.Data.Name)
	}
	if cmd.Data.Description != nil {
		category.Description = strings.TrimSpace(*cmd.Data.Description)
	}
	if cmd.Data.Status != nil {
		if *cmd.Data.Status < 0 || *cmd.Data.Status > 1 {
			return  errors.New(categorydomain.ErrInvalidStatusFilter)
		}
		if category.Status == 0 {
			return errors.New(categorydomain.ErrCategoryInactive)
		}
		category.Status = *cmd.Data.Status
	}

	return s.repo.Update(ctx, category, &cmd.Id)
}