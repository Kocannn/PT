package services

import "github.com/kocannn/api-programming-tadulako/domain"

type CategoriesService struct {
	CategoriesRepo domain.CategoriesRepository
}

// Delete implements domain.CategoriesService.
func (c *CategoriesService) Delete(Id int) error {
	return c.CategoriesRepo.Delete(Id)
}

func (c *CategoriesService) Create(categories *domain.Categories) (*domain.Categories, error) {
	return c.CategoriesRepo.Create(categories)
}

func NewCategoriesService(categoriesRepo domain.CategoriesRepository) domain.CategoriesService {
	return &CategoriesService{
		CategoriesRepo: categoriesRepo,
	}
}
