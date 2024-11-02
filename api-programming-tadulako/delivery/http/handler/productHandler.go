package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kocannn/api-programming-tadulako/domain"
	"github.com/kocannn/api-programming-tadulako/dto"
)

func NewProductHandler(productServices domain.ProductsService) *ProductHandler {
	return &ProductHandler{
		ProductService: productServices,
	}
}

type ProductHandler struct {
	ProductService domain.ProductsService
}

func (h *ProductHandler) Create(c *gin.Context) {
	var Product domain.Products

	role, _ := c.Get("role")

	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Only admins can create products",
		})
		return
	}

	if err := c.ShouldBindJSON(&Product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	product, err := h.ProductService.Create(&Product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"product": product,
	})
}

func (h *ProductHandler) GetById(c *gin.Context) {
	id := c.Param("id")

	product, err := h.ProductService.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})

}

func (h *ProductHandler) Update(c *gin.Context) {
	id := c.Param("id")

	role, _ := c.Get("role")

	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Only admins can update products",
		})
		return
	}

	product, err := h.ProductService.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := h.ProductService.Update(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
	})

}
func (h *ProductHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	role, _ := c.Get("role")

	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Only admins can delete products",
		})
		return
	}

	if err := h.ProductService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(), "status": http.StatusInternalServerError,
			"message": "Failed to delete product",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}
func (h *ProductHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.Param("page"))

	products, err := h.ProductService.List(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(products) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No products found",
		})
		return
	}

	productDTOs := []*dto.ProductDTO{}
	for _, product := range products {
		productDTO := &dto.ProductDTO{
			ID:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
		}
		for _, category := range product.Categories {
			categoryDTO := &dto.CategoryDTO{
				Id:          category.Id,
				Name:        category.Name,
				Description: category.Description,
			}
			productDTO.Category = append(productDTO.Category, *categoryDTO)
		}
		productDTOs = append(productDTOs, productDTO)
	}

	c.JSON(http.StatusOK, gin.H{
		"products": productDTOs,
	})
}

func (h *ProductHandler) AddCategories(c *gin.Context) {
	role, _ := c.Get("role")

	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Only admins can add categories to products",
		})
		return
	}

	var category *dto.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := h.ProductService.AddCategories(category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product categories added successfully",
	})
}
