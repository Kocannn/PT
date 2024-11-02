package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kocannn/api-programming-tadulako/dto"
)

type Products struct {
	Id          int           `json:"id" gorm:"primary_key:auto_increment"`
	Name        string        `json:"name" gorm:"type:varchar(50);not null"`
	Description string        `json:"description" gorm:"type:text"`
	Price       float64       `json:"price" gorm:"type:decimal(10,2);not null"`
	Stock       int           `json:"stock" gorm:"type:int;not null"`
	Categories  []Categories  `json:"categories,omitempty" gorm:"many2many:products_categories;constraint:OnDelete:CASCADE"`
	Order_items []Order_items `json:"order_items,omitempty" gorm:"foreignKey:Order_id;references:Id"`
	CreatedAt   time.Time     `json:"created_at" gorm:"type:datetime;autoCreateTime"`
	UpdatedAt   time.Time     `json:"updated_at" gorm:"type:datetime;autoUpdateTime"`
}

type ProductsRepository interface {
	Create(product *Products) (*Products, error)
	GetById(id string) (*Products, error)
	Update(product *Products) error
	Delete(id string) error
	List(limit, offsets int) ([]*Products, error)
	AddCategories(product_id, category_id []int) error
}

type ProductsService interface {
	Create(product *Products) (*Products, error)
	GetById(id string) (*Products, error)
	Update(product *Products) error
	Delete(id string) error
	List(limit int) ([]*Products, error)
	AddCategories(*dto.Category) error
}

type ProductHandler interface {
	Create(c *gin.Context)
	GetById(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	List(c *gin.Context)
	AddCategories(c *gin.Context)
}
