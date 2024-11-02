package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kocannn/api-programming-tadulako/delivery/http/handler"
	"github.com/kocannn/api-programming-tadulako/delivery/http/middleware"
	"github.com/kocannn/api-programming-tadulako/repository"
	"github.com/kocannn/api-programming-tadulako/services"
	"gorm.io/gorm"
)

func UserRoutes(group *gin.RouterGroup, db *gorm.DB) {

	userRepo := repository.NewUserRepo(db)
	userServices := services.NewUserServices(userRepo)
	userHandler := handler.NewUserHandler(userServices)
	group.POST("/register", userHandler.Register)
	group.POST("/login", userHandler.Login)

	auth := group.Use(middleware.AuthMiddleware())
	auth.GET("/validate", middleware.AuthMiddleware(), func(c *gin.Context) {
		user, err := c.Get("id")
		if !err {
			return
		}
		role, _ := c.Get("role")
		c.JSON(http.StatusOK, gin.H{
			"message": "Logging in successfully",
			"user":    user,
			"role":    role,
		})
	})

}
