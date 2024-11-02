package repository

import (
	"errors"
	"log"

	"github.com/kocannn/api-programming-tadulako/domain"
	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

// UpdateStatusOrder implements domain.OrderRepository.
func (o *orderRepo) UpdateStatusOrder(id int, order *domain.Orders) error {
	if err := o.db.Model(&domain.Orders{}).Where("id = ?", id).Update("status", order.Status).Error; err != nil {
		return err
	}
	if order.Status == "" {
		return errors.New("invalid order status")
	}
	return nil
}

// Create implements domain.OrderRepository.
func (o *orderRepo) Create(order *domain.Orders, items []*domain.Order_items) error {

	if err := o.db.Create(order).Error; err != nil {
		return err
	}
	log.Println(items)
	for _, item := range items {
		item.Order_id = order.Id

		log.Println(*item)
		if err := o.db.Model(&item).Create(&item).Error; err != nil {
			return err
		}
	}

	if err := o.db.Model(&order).Association("Order_items").Append(items); err != nil {
		return err
	}

	return nil
}

// GetById implements domain.OrderRepository.
func (o *orderRepo) GetById(id int) (*domain.Orders, error) {
	product := &domain.Orders{}
	if err := o.db.Preload("Order_items.Products").First(&product, id).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// GetByUserId implements domain.OrderRepository.
func (o *orderRepo) GetByUserId(id int) ([]*domain.Orders, error) {
	var domainOrders []*domain.Orders
	if err := o.db.Preload("Order_items.Products").Where("user_id = ?", id).Find(&domainOrders).Error; err != nil {
		return nil, err
	}
	return domainOrders, nil
}

func (m *orderRepo) List(limit, offsets, id int) ([]*domain.Orders, error) {
	order := []*domain.Orders{}
	if err := m.db.Preload("Order_items.Products").Limit(limit).Offset(offsets).Where("user_id = ?", id).Find(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return order, nil
		}
		return order, err
	}
	return order, nil
}

func NewOrderRepo(db *gorm.DB) domain.OrderRepository {
	return &orderRepo{
		db: db,
	}
}
