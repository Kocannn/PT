package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kocannn/api-programming-tadulako/delivery/http/handler"
	"github.com/kocannn/api-programming-tadulako/delivery/http/middleware"
	"github.com/kocannn/api-programming-tadulako/repository"
	"github.com/kocannn/api-programming-tadulako/services"
	"gorm.io/gorm"
)

func ProductRoutes(group *gin.RouterGroup, db *gorm.DB) {
	productRepo := repository.NewProductRepo(db)
	productService := services.NewProductServices(productRepo)
	productHandler := handler.NewProductHandler(productService)

	group.GET("/:id", productHandler.GetById)
	group.GET("/page/:page", productHandler.List)
	auth := group.Use(middleware.AuthMiddleware())
	auth.POST("/addCategories/", productHandler.AddCategories)
	auth.DELETE("/:id", productHandler.Delete)
	auth.PUT("/:id", productHandler.Update)
	auth.POST("/create", productHandler.Create)
}
