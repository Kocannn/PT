package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kocannn/api-programming-tadulako/delivery/http/handler"
	"github.com/kocannn/api-programming-tadulako/repository"
	"github.com/kocannn/api-programming-tadulako/services"

	"gorm.io/gorm"
)

func AdminRepo(group *gin.RouterGroup, db *gorm.DB) {
	adminRepo := repository.NewAdminRepo(db)
	orderRepo := repository.NewOrderRepo(db)
	adminService := services.NewAdminService(adminRepo, orderRepo)
	adminHandler := handler.NewAdminHandler(adminService)

	group.POST("/login", adminHandler.Login)
	group.POST("/register", adminHandler.Create)
	group.PUT("/:id", adminHandler.Update)

}
