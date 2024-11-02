package repository

import (
	"errors"

	"github.com/kocannn/api-programming-tadulako/domain"
	"gorm.io/gorm"
)

type AdminRepo struct {
	db *gorm.DB
}

// GetByEmail implements domain.AdminRepository.
func (a *AdminRepo) GetByEmail(email string) (*domain.Admin, error) {
	admin := &domain.Admin{}
	result := a.db.First(admin, "email = ?", email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return admin, nil
		}
		return admin, result.Error
	}
	return admin, nil
}

func NewAdminRepo(db *gorm.DB) domain.AdminRepository {
	return &AdminRepo{
		db: db,
	}
}

// Create implements domain.AdminRepository.
func (a *AdminRepo) Create(admin *domain.Admin) error {

	if err := a.db.Create(admin).Error; err != nil {
		return err
	}
	return nil

}
