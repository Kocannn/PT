package repository

import (
	"github.com/kocannn/api-programming-tadulako/domain"
	"gorm.io/gorm"
)

type categoriesRepo struct {
	db *gorm.DB
}

func NewCategoriesRepo(db *gorm.DB) domain.CategoriesRepository {
	return &categoriesRepo{
		db: db,
	}
}

func (m *categoriesRepo) Create(categories *domain.Categories) (*domain.Categories, error) {
	if err := m.db.Create(categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (m *categoriesRepo) Delete(Id int) error {
	if err := m.db.Where("id = ?", Id).Delete(&domain.Categories{}).Error; err != nil {
		return err
	}
	return nil
}
