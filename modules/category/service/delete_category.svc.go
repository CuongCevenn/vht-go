package categoryservice

import (
	"context"
	categorydtos "vht-go/modules/category/dtos"

	"github.com/google/uuid"
)

type DeleteCategoryCommand struct {
	Id *uuid.UUID
}

type IDeleteCategoryRepository interface {
	Delete(ctx context.Context, id *uuid.UUID) error
}

type DeleteCategoryCommandHandler struct {
	queryRepo IGetCategoryQueryRepository
	repo IDeleteCategoryRepository
}

func NewDeleteCategoryCommandHandler(queryRepo IGetCategoryQueryRepository, repo IDeleteCategoryRepository) *DeleteCategoryCommandHandler {
	return &DeleteCategoryCommandHandler{queryRepo: queryRepo, repo: repo}
}

func (h *DeleteCategoryCommandHandler) Handle(ctx context.Context, cmd *DeleteCategoryCommand) error {
	_, err := h.queryRepo.FindById(ctx, &categorydtos.GetCategoryDTO{Id: cmd.Id})
	if err != nil {
		return err
	}
	return h.repo.Delete(ctx, cmd.Id)
}