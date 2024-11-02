package domain

type Order_items struct {
	Id          int      `json:"id" gorm:"primary_key auto_increment"`
	Order_id    int      `json:"order_id" gorm:"type:int;not null"`
	Product_id  int      `json:"product_id" gorm:"type:int;not null"`
	NameProduct string   `json:"product_name,omitempty" gorm:"type:varchar(50);not null"`
	Quantity    int      `json:"quantiy" gorm:"type:int;not null"`
	Price       float64  `json:"price" gorm:"type:decimal(10,2);not null"`
	Orders      Orders   `json:"-" gorm:"foreignKey:Order_id;references:Id"`
	Products    Products `json:"-" gorm:"foreignKey:Product_id;references:Id"`
}
