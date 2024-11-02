package services

import (
	"errors"
	"strconv"

	"github.com/kocannn/api-programming-tadulako/domain"
	"github.com/kocannn/api-programming-tadulako/dto"
)

type OrdersService struct {
	OrderRepo   domain.OrderRepository
	ProductRepo domain.ProductsRepository
}

// Update implements domain.OrderService.

// CreateOrder implements domain.OrderService.
func (o *OrdersService) CreateOrder(order *domain.Orders, items []*dto.InputOrder) error {
	if order.User_id == 0 || len(items) == 0 {
		return errors.New("invalid order input")
	}
	var orderItems []*domain.Order_items
	var totalAmount float64
	for _, item := range items {
		product, err := o.ProductRepo.GetById(strconv.Itoa(item.ProductId))
		if err != nil {
			return err
		}
		if product.Stock < item.Quantity {
			return errors.New("insufficient stock")
		}
		totalAmount += product.Price * float64(item.Quantity)
		product.Stock -= item.Quantity
		o.ProductRepo.Update(product)
		orderItems = append(orderItems, &domain.Order_items{
			Order_id:    order.Id,
			Product_id:  item.ProductId,
			NameProduct: product.Name,
			Quantity:    item.Quantity,
			Price:       product.Price,
		})
	}
	order.Total_amount = totalAmount

	if err := o.OrderRepo.Create(order, orderItems); err != nil {

		return err
	}

	return nil
}

// GetOrderById implements domain.OrderService.
func (o *OrdersService) GetOrderById(id int) (*domain.Orders, error) {
	order, err := o.OrderRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return order, nil

}

// GetOrderByUserId implements domain.OrderService.
func (o *OrdersService) GetOrderByUserId(id int) ([]*domain.Orders, error) {
	order, err := o.OrderRepo.GetByUserId(id)
	if err != nil {
		return nil, err
	}
	if len(order) == 0 {
		return nil, errors.New("order not found")
	}

	return order, nil

}

func (o *OrdersService) List(page, id int) ([]*domain.Orders, error) {
	limit := 2
	offset := (page - 1) * limit
	return o.OrderRepo.List(limit, offset, id)
}

func NewOrderServices(or domain.OrderRepository, pr domain.ProductsRepository) domain.OrderService {
	return &OrdersService{
		OrderRepo:   or,
		ProductRepo: pr,
	}
}
