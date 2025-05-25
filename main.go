package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kuldeep/golang-jwt-project/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	r := gin.New()
	r.Use(gin.Logger())

	router.UserRouts(r)
	router.AuthRoutes(r)

	r.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted to api-1"})
	})

	r.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted to api-2"})
	})

	r.Run(":" + port)
}
