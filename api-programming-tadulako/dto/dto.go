package dto

type InputOrder struct {
	ProductId int `json:"product_id"`
	OrderId   int `json:"order_id"`
	Quantity  int `json:"quantity"`
}

type OrderDTO struct {
	ID          int            `json:"id"`
	UserID      int            `json:"user_id"`
	TotalAmount float64        `json:"total_amount"`
	Status      string         `json:"status,omitempty"`
	OrderItems  []OrderItemDTO `json:"order_items,omitempty"`
}

type OrderItemDTO struct {
	ID          int         `json:"id"`
	OrderID     int         `json:"order_id"`
	ProductName string      `json:"product_name,omitempty"`
	ProductID   int         `json:"product_id"`
	Quantity    int         `json:"quantiy"`
	Price       float64     `json:"price"`
	Products    *ProductDTO `json:"products,omitempty"`
}

type NewProduct struct {
	Id           int               `json:"id"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	Price        float64           `json:"price"`
	Stock        int               `json:"stock"`
	Categories   []NewCategories   `json:"categories"`
	OrderItemDTO []NewOrderItemDTO `json:"orderItemDTOList"`
}

type NewOrderItemDTO struct {
	ID        int     `json:"id"`
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantiy"`
	Price     float64 `json:"price"`
}

type NewCategories struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Category struct {
	ProductId  []int `json:"product_id"`
	CategoryId []int `json:"category_id"`
}

type ProductDTO struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description,omitempty"`
	Price       float64       `json:"price"`
	Stock       int           `json:"stock"`
	Category    []CategoryDTO `json:"categories"`
}

type CategoryDTO struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
