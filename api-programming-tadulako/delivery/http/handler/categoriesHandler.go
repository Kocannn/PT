package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kocannn/api-programming-tadulako/domain"
)

type CategoriesHandler struct {
	CategoriesService domain.CategoriesService
}

func NewCategoriesHandler(categoriesService domain.CategoriesService) *CategoriesHandler {
	return &CategoriesHandler{
		CategoriesService: categoriesService,
	}
}

func (h *CategoriesHandler) Create(c *gin.Context) {
	var inputCategories domain.Categories
	if err := c.ShouldBindJSON(&inputCategories); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	categories, err := h.CategoriesService.Create(&inputCategories)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully", "data": categories})

}
func (h *CategoriesHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.CategoriesService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})

}
