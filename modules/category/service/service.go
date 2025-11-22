package categoryservice

import categoryrepository "vht-go/modules/category/repository"

type CategoryService struct {
	repo *categoryrepository.GORMCategoryRepository
}

func NewCategoryService(repo *categoryrepository.GORMCategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}