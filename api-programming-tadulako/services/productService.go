package services

import (
	"github.com/kocannn/api-programming-tadulako/domain"
	"github.com/kocannn/api-programming-tadulako/dto"
)

type ProductsService struct {
	ProductRepo domain.ProductsRepository
}

func NewProductServices(productRepo domain.ProductsRepository) domain.ProductsService {
	return &ProductsService{
		ProductRepo: productRepo,
	}
}

func (p *ProductsService) Delete(id string) error {
	return p.ProductRepo.Delete(id)
}

func (p *ProductsService) GetById(id string) (*domain.Products, error) {
	product, err := p.ProductRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return product, nil

}

func (p *ProductsService) List(page int) ([]*domain.Products, error) {

	limit := 10
	offset := (page - 1) * limit
	return p.ProductRepo.List(limit, offset)
}

func (p *ProductsService) Update(product *domain.Products) error {
	return p.ProductRepo.Update(product)
}

func (p *ProductsService) Create(product *domain.Products) (*domain.Products, error) {
	return p.ProductRepo.Create(product)
}

func (p *ProductsService) AddCategories(category *dto.Category) error {
	return p.ProductRepo.AddCategories(category.ProductId, category.CategoryId)
}
