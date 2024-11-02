package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kocannn/api-programming-tadulako/dto"
)

type Orders struct {
	Id           int           `json:"id" gorm:"primary_key auto_increment"`
	User_id      int           `json:"user_id" gorm:"type:int;not null" `
	Total_amount float64       `json:"total_amount" gorm:"type:decimal(10,2);not null"`
	Status       string        `json:"status" gorm:"type:enum('pending', 'paid', 'shipped', 'delivered', 'cancelled');default:'pending';not null"`
	Order_items  []Order_items `json:"-" gorm:"foreignKey:Order_id"`
	CreatedAt    time.Time     `json:"created_at" gorm:"type:datetime;autoCreateTime"`
	UpdatedAt    time.Time     `json:"updated_at" gorm:"type:datetime;autoUpdateTime"`
}

type OrderHandler interface {
	CreateOrder(c *gin.Context)
	GetOrder(c *gin.Context)
	GetUserOrder(c *gin.Context)
	List(c *gin.Context)
}

type OrderRepository interface {
	Create(order *Orders, items []*Order_items) error
	GetById(id int) (*Orders, error)
	GetByUserId(id int) ([]*Orders, error)
	UpdateStatusOrder(id int, order *Orders) error
	List(page, offsets, id int) ([]*Orders, error)
}

type OrderService interface {
	CreateOrder(order *Orders, items []*dto.InputOrder) error
	GetOrderById(id int) (*Orders, error)
	GetOrderByUserId(id int) ([]*Orders, error)
	List(page, id int) ([]*Orders, error)
}
