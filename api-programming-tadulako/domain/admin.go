package domain

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Admin struct {
	Id        int       `json:"id" gorm:"primary_key auto_increment"`
	Username  string    `json:"username" gorm:"type:varchar(50);unique;not null"`
	Email     string    `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	Role      string    `json:"role" gorm:"type:enum('user', 'admin');default:'admin'"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime;autoUpdateTime"`
}

type AdminRepository interface {
	Create(admin *Admin) error
	GetByEmail(email string) (*Admin, error)
}

type AdminService interface {
	Create(admin *Admin) error
	GetById(id int) (*Orders, error)
	Update(id int, order *Orders) error
	Login(email, password string) (string, error)
}

type AdminHandler interface {
	Create(c *gin.Context)
	GetById(c *gin.Context)
	Login(c *gin.Context)
}
