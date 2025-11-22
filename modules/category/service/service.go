package categoryservice

import categoryrepository "vht-go/modules/category/repository"

type CategoryService struct {
	repo *categoryrepository.GORMCategoryRepository
}