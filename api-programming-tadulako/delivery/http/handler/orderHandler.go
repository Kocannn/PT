package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kocannn/api-programming-tadulako/domain"
	"github.com/kocannn/api-programming-tadulako/dto"
)

type OrderHandler struct {
	OrderService domain.OrderService
}

func NewOrderHandler(orderService domain.OrderService) *OrderHandler {
	return &OrderHandler{
		OrderService: orderService,
	}
}

func (h *OrderHandler) GetOrderById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	order, err := h.OrderService.GetOrderById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	orderDTO := &dto.OrderDTO{
		ID:          order.Id,
		UserID:      order.User_id,
		TotalAmount: order.Total_amount,
		Status:      order.Status,
	}

	for _, item := range order.Order_items {
		orderItemDTO := dto.OrderItemDTO{
			ID:        item.Id,
			OrderID:   item.Order_id,
			ProductID: item.Product_id,
			Quantity:  item.Quantity,
			Price:     item.Price,
			Products: &dto.ProductDTO{
				ID:          item.Products.Id,
				Name:        item.Products.Name,
				Description: item.Products.Description,
				Price:       item.Products.Price,
				Stock:       item.Products.Stock,
			},
		}
		orderDTO.OrderItems = append(orderDTO.OrderItems, orderItemDTO)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": orderDTO,
	})
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var InputOrder []*dto.InputOrder

	role, _ := c.Get("role")
	if role != "user" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Only users can create orders",
		})
		return
	}

	if err := c.ShouldBindJSON(&InputOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	userId, _ := c.Get("id")
	userIdInt := userId.(int)
	order := &domain.Orders{
		User_id: userIdInt,
	}

	if err := h.OrderService.CreateOrder(order, InputOrder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	orderDTO := &dto.OrderDTO{
		ID:          order.Id,
		UserID:      order.User_id,
		TotalAmount: order.Total_amount,
		Status:      order.Status,
	}

	for _, item := range order.Order_items {
		orderItemDTO := dto.OrderItemDTO{
			ID:          item.Id,
			OrderID:     item.Order_id,
			ProductID:   item.Product_id,
			ProductName: item.NameProduct,
			Quantity:    item.Quantity,
			Price:       item.Price,
		}
		orderDTO.OrderItems = append(orderDTO.OrderItems, orderItemDTO)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order created successfully",
		"data":    orderDTO,
	})
}

func (h *OrderHandler) GetByUserId(c *gin.Context) {
	userId, _ := c.Get("id")
	userIdInt := userId.(int)

	orders, err := h.OrderService.GetOrderByUserId(userIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var ordersDTO []*dto.OrderDTO
	for _, order := range orders {
		orderDTO := &dto.OrderDTO{
			ID:          order.Id,
			UserID:      order.User_id,
			TotalAmount: order.Total_amount,
			Status:      order.Status,
		}

		for _, item := range order.Order_items {
			orderItemDTO := dto.OrderItemDTO{
				ID:          item.Id,
				OrderID:     item.Order_id,
				ProductID:   item.Product_id,
				ProductName: item.Products.Name,
				Quantity:    item.Quantity,
				Price:       item.Price,
			}
			orderDTO.OrderItems = append(orderDTO.OrderItems, orderItemDTO)
		}

		ordersDTO = append(ordersDTO, orderDTO)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ordersDTO,
	})
}
func (h *OrderHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.Param("page"))
	userId, _ := c.Get("id")
	userIdInt := userId.(int)

	orders, err := h.OrderService.List(page, userIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var ordersDTO []*dto.OrderDTO
	for _, order := range orders {
		orderDTO := &dto.OrderDTO{
			ID:          order.Id,
			UserID:      order.User_id,
			TotalAmount: order.Total_amount,
			Status:      order.Status,
		}

		for _, item := range order.Order_items {
			orderItemDTO := dto.OrderItemDTO{
				ID:        item.Id,
				OrderID:   item.Order_id,
				ProductID: item.Product_id,
				Quantity:  item.Quantity,
				Price:     item.Price,
			}
			orderDTO.OrderItems = append(orderDTO.OrderItems, orderItemDTO)
		}

		ordersDTO = append(ordersDTO, orderDTO)
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": ordersDTO,
	})
}
