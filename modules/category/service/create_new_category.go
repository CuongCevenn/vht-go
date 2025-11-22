package categoryservice

import (
	"context"
	"errors"
	"strings"
	"time"
	categorydomain "vht-go/modules/category/domain"

	"github.com/google/uuid"
)

type CreateCategoryDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (dto *CreateCategoryDTO) Validate() error {
	dto.Name = strings.TrimSpace(dto.Name)
	dto.Description = strings.TrimSpace(dto.Description)

	if dto.Name == "" {
		return errors.New(categorydomain.ErrCategoryNameRequired)
	}
	return nil
}

func (s *CategoryService) CreateNewCategory(ctx context.Context, dto *CreateCategoryDTO) (newId *uuid.UUID, err error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	newCatId, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	category := categorydomain.Category{
		Id: 		newCatId,
		Name: 	  dto.Name,
		Description: dto.Description,
		Status:    1,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := s.repo.Insert(ctx, &category); err != nil {
		return nil, err
	}
	
	return &newCatId, nil
}