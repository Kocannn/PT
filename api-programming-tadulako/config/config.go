package config

import (
	"fmt"
	"log"
	"os"

	"github.com/kocannn/api-programming-tadulako/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDB() *gorm.DB {

	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_host := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", db_user, db_pass, db_host)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&domain.User{}, &domain.Products{}, &domain.Categories{}, &domain.Orders{}, &domain.Order_items{}, &domain.Admin{})

	log.Println("database is connected")

	return db
}
