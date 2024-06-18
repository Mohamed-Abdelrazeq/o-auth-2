package main

import (
	"Mohamed-Abdelrazeq/o-auth-2/handlers"
	"Mohamed-Abdelrazeq/o-auth-2/loaders"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load Environment Variables
	loaders.LoadEnv()

	// Create Router
	r := gin.Default()

	// Test Router
	r.GET("/ping", handlers.Ping)

	// Run Server
	r.Run()
}
