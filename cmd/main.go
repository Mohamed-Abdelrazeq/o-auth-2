package main

import (
	apps "Mohamed-Abdelrazeq/o-auth-2/internal/application"
	"Mohamed-Abdelrazeq/o-auth-2/internal/handlers"
	"Mohamed-Abdelrazeq/o-auth-2/internal/loaders"
	"Mohamed-Abdelrazeq/o-auth-2/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load Env & DB
	loaders.LoadEnv()
	db := loaders.LoadDB()

	// Init AuthService
	authService := services.NewAuthSericeInstance(db)
	authApp := apps.NewAuthApplicationInstance(authService)
	authHandler := handlers.NewAuthHandlerInstance(authApp)

	// Create Router
	r := gin.Default()

	// Test Router
	r.GET("/ping", handlers.Ping)

	// Auth Router
	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)

	// Run Server
	r.Run()
}
