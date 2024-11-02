package repository

import (
	"errors"
	"fmt"

	"github.com/kocannn/api-programming-tadulako/domain"
	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) domain.ProductsRepository {
	return &productRepo{
		db: db,
	}
}

func (m *productRepo) Create(productRepo *domain.Products) (*domain.Products, error) {
	if err := m.db.Create(productRepo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	return productRepo, nil
}

func (m *productRepo) GetById(id string) (*domain.Products, error) {
	product := &domain.Products{}

	err := m.db.Preload("Order_items").Preload("Categories").First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (m *productRepo) Update(productRepo *domain.Products) error {
	if err := m.db.Save(productRepo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	return nil
}

func (m *productRepo) Delete(id string) error {
	product := &domain.Products{}
	if err := m.db.Delete(product, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	return nil
}

func (m *productRepo) List(limit, offsets int) ([]*domain.Products, error) {
	product := []*domain.Products{}
	if err := m.db.Limit(limit).Offset(offsets).Preload("Categories").Find(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, nil
		}
		return product, err
	}
	return product, nil
}

func (m *productRepo) AddCategories(product_id, category []int) error {
	for _, productId := range product_id {
		var product domain.Products
		if err := m.db.First(&product, productId).Error; err != nil {
			return fmt.Errorf("product not found with ID %d: %w", productId, err)
		}

		for _, categoryId := range category {
			var category domain.Categories
			if err := m.db.First(&category, categoryId).Error; err != nil {
				return fmt.Errorf("category not found with ID %d: %w", categoryId, err)
			}

			if err := m.db.Model(&product).Association("Categories").Append(&category); err != nil {
				return fmt.Errorf("failed to associate category %d with product %d: %w", categoryId, productId, err)
			}
		}
	}
	return nil
}
