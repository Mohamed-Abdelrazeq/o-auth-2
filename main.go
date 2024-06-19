package main

import (
	"Mohamed-Abdelrazeq/o-auth-2/handlers"
	"Mohamed-Abdelrazeq/o-auth-2/loaders"
	"Mohamed-Abdelrazeq/o-auth-2/models"
	"Mohamed-Abdelrazeq/o-auth-2/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load Env & DB
	loaders.LoadEnv()
	db := loaders.LoadDB()

	// Init AuthService
	authService := services.NewAuthSericeInstance(db)

	// Create Router
	r := gin.Default()

	// Test Router
	r.GET("/ping", handlers.Ping)

	// Auth Router
	r.GET("/login", func(ctx *gin.Context) {
		user, err := authService.GetUser("mohamed@gmail.com")
		if err != nil {
			fmt.Println(err)
		}
		ctx.JSON(200, user)
	})

	r.POST("/register", func(ctx *gin.Context) {
		user := models.User{
			Email:    "mohamed@gmail.com",
			Password: "123456",
		}
		err := authService.CreateUser(&user)
		if err != nil {
			fmt.Println(err)
		}
	})

	// Run Server
	r.Run()
}
