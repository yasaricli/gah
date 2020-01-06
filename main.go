package main

import (
	"./handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", handlers.LoginHandler)
			auth.POST("/register", handlers.RegisterHandler)
		}
	}

	router.Run(":4000")
}
