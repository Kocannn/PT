package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kocannn/api-programming-tadulako/domain"
)

type AdminHandler struct {
	AdminService domain.AdminService
}

func NewAdminHandler(adminService domain.AdminService) *AdminHandler {
	return &AdminHandler{
		AdminService: adminService,
	}
}

func (h *AdminHandler) Login(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := h.AdminService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *AdminHandler) Create(c *gin.Context) {
	var admin domain.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.AdminService.Create(&admin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Admin created successfully", "data": admin})
}

func (h *AdminHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	type data struct {
		Status string `json:"status"`
	}
	var input data
	order, err := h.AdminService.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.Status = input.Status

	if err := h.AdminService.Update(id, order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully", "data": order})

}
