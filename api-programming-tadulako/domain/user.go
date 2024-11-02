package domain

import (
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id        int       `json:"id" gorm:"primary_key auto_increment"`
	Username  string    `json:"username" gorm:"type:varchar(50);unique;not null"`
	Email     string    `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	Full_name string    `json:"fullname" gorm:"type:varchar(100)"`
	Address   string    `json:"address" gorm:"type:text"`
	Phone     string    `json:"Phone" gorm:"type:varchar(20)"`
	Role      string    `json:"role" gorm:"type:enum('user', 'admin');default:'user'"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime;autoUpdateTime"`
	Orders    []Orders  `gorm:"foreignKey:User_id"`
}

type UsersRepository interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
}

type UserServices interface {
	Register(user *User) error
	Login(email, password string) (string, error)
}

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}
