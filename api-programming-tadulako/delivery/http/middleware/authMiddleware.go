package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kocannn/api-programming-tadulako/domain"
	"gorm.io/gorm"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authCookie, err := c.Cookie("Authorization")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authCookie, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			secret := []byte(os.Getenv("SECRET"))
			return secret, nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		db, exists := c.Get("db")
		if !exists || db == nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		email := claims["email"].(string)
		user, _ := getUserByEmail(email, db.(*gorm.DB))
		admin, _ := getAdminByEmail(email, db.(*gorm.DB))

		if user != nil {
			c.Set("id", user.Id)
			c.Set("role", user.Role)
			c.Set("user", user)
			c.Next()
		}
		if admin != nil {
			c.Set("id", admin.Id)
			c.Set("role", admin.Role)
			c.Set("admin", admin)
			c.Next()
		}

	}
}

func getAdminByEmail(email string, db *gorm.DB) (*domain.Admin, error) {
	var admin domain.Admin

	if err := db.Where("email = ?", email).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &admin, nil
}

func getUserByEmail(email string, db *gorm.DB) (*domain.User, error) {
	var user domain.User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
