package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kocannn/api-programming-tadulako/config"
	routes "github.com/kocannn/api-programming-tadulako/delivery/http/routes"
)



func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	

	r := gin.Default()
	
	db := config.ConnectionDB()
	
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
})

	routes.Routes(r, db)

  

  r.Run("localhost:8080");
}