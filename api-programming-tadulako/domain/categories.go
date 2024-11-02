package domain

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Categories struct {
	Id          int        `json:"id" gorm:"primary_key auto_increment"`
	Name        string     `json:"name" gorm:"type:varchar(50);unique;not null"`
	Description string     `json:"description" gorm:"type:text"`
	Products    []Products `json:"products,omitempty" gorm:"many2many:products_categories"`
	CreatedAt   time.Time  `json:"created_at" gorm:"type:datetime;autoCreateTime"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"type:datetime;autoUpdateTime"`
}

type CategoriesHandler interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
}

type CategoriesService interface {
	Create(Categories *Categories) (*Categories, error)
	Delete(Id int) error
}

type CategoriesRepository interface {
	Create(Categories *Categories) (*Categories, error)
	Delete(Id int) error
}
