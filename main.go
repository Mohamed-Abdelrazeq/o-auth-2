package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load Environment Variables
	loadEnv()

	// Create Router
	r := gin.Default()

	// Test Router
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(
			200,
			gin.H{
				"message": "pong",
			})
	})

	// Run Server
	r.Run()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
