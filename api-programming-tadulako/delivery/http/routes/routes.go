package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(r *gin.Engine, db *gorm.DB) {
	user := r.Group("/api/v1/users")
	product := r.Group("/api/v1/products")
	order := r.Group("/api/v1/orders")
	admin := r.Group("/api/v1/admin")
	categories := r.Group("/api/v1/categories")
	CategoriesRoutes(categories, db)
	OrderRoutes(order, db)
	UserRoutes(user, db)
	ProductRoutes(product, db)
	AdminRepo(admin, db)

}
