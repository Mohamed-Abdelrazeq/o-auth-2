package main

import (
	apps "Mohamed-Abdelrazeq/o-auth-2/internal/application"
	"Mohamed-Abdelrazeq/o-auth-2/internal/handlers"
	"Mohamed-Abdelrazeq/o-auth-2/internal/loaders"
	"Mohamed-Abdelrazeq/o-auth-2/internal/middlewares"
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

	// Test Routes
	r.GET("/ping", handlers.HealthTest)

	// Auth Routes
	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)

	// Protected Routes
	r.Use(middlewares.AuthenticateUser())
	r.GET("/protected", handlers.AuthenticationTest)

	// Run Server
	r.Run()
}
