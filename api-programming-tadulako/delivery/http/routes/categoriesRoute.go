package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kocannn/api-programming-tadulako/delivery/http/handler"
	"github.com/kocannn/api-programming-tadulako/delivery/http/middleware"
	"github.com/kocannn/api-programming-tadulako/repository"
	"github.com/kocannn/api-programming-tadulako/services"
	"gorm.io/gorm"
)

func CategoriesRoutes(group *gin.RouterGroup, db *gorm.DB) {
	categoriesRepo := repository.NewCategoriesRepo(db)
	categoriesService := services.NewCategoriesService(categoriesRepo)
	categoriesHandler := handler.NewCategoriesHandler(categoriesService)

	auth := group.Use(middleware.AuthMiddleware())
	auth.POST("/create", categoriesHandler.Create)
	auth.DELETE("/delete/:id", categoriesHandler.Delete)

}
