package categoryrepository

import (
	"context"
	"errors"
	categorydomain "vht-go/modules/category/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (repo *GORMCategoryRepository) FindAll(ctx context.Context, paging *categorydomain.Paging) (categories []categorydomain.Category, err error) {
	if err = repo.db.Table(categorydomain.Category{}.TableName()).Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).Order("id desc").Count(&paging.Total).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}	

func (repo *GORMCategoryRepository) FindById(ctx context.Context, id *uuid.UUID) (category *categorydomain.Category, err error) {
	if err = repo.db.First(&category, "id = ?", *id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New(categorydomain.ErrCategoryNotFound)
		}
		return nil, err
	}
	return category, nil
}