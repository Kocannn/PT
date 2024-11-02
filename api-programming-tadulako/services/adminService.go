package services

import (
	"errors"

	"github.com/kocannn/api-programming-tadulako/domain"
	"github.com/kocannn/api-programming-tadulako/helper"
)

type AdminService struct {
	AdminRepository domain.AdminRepository
	OrderRepository domain.OrderRepository
}

// Login implements domain.AdminService.
func (a *AdminService) Login(email string, password string) (string, error) {
	admin, err := a.AdminRepository.GetByEmail(email)
	if err != nil {
		return "", err
	}

	isPasswordValid, err := helper.ComparePass([]byte(admin.Password), []byte(password))
	if err != nil {
		return "", err
	}
	if !isPasswordValid {
		return "", errors.New("invalid password")
	}

	token, err := generateJWTToken(nil, admin)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetById implements domain.AdminService.
func (a *AdminService) GetById(id int) (*domain.Orders, error) {
	return a.OrderRepository.GetById(id)
}

func NewAdminService(adminRepository domain.AdminRepository, orderRepository domain.OrderRepository) domain.AdminService {
	return &AdminService{
		AdminRepository: adminRepository,
		OrderRepository: orderRepository,
	}
}

// Create implements domain.AdminRepository.
func (a *AdminService) Create(admin *domain.Admin) error {
	hassPass, err := helper.HassPass(admin.Password)
	if err != nil {
		return err
	}
	admin.Password = hassPass

	return a.AdminRepository.Create(admin)
}

func (a *AdminService) Update(id int, order *domain.Orders) error {
	return a.OrderRepository.UpdateStatusOrder(id, order)
}
