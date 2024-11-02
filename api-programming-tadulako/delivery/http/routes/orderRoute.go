package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kocannn/api-programming-tadulako/delivery/http/handler"
	"github.com/kocannn/api-programming-tadulako/delivery/http/middleware"
	"github.com/kocannn/api-programming-tadulako/repository"
	"github.com/kocannn/api-programming-tadulako/services"
	"gorm.io/gorm"
)

func OrderRoutes(group *gin.RouterGroup, db *gorm.DB) {
	orderRepo := repository.NewOrderRepo(db)
	productRepo := repository.NewProductRepo(db)
	orderService := services.NewOrderServices(orderRepo, productRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	auth := group.Use(middleware.AuthMiddleware())
	auth.POST("/create", orderHandler.CreateOrder)
	auth.GET("/:id", orderHandler.GetOrderById)
	auth.GET("/", orderHandler.GetByUserId)
	auth.GET("/page/:page", orderHandler.List)

}
