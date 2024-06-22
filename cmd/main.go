package main

import (
	"Mohamed-Abdelrazeq/o-auth-2/internal/database"
	"Mohamed-Abdelrazeq/o-auth-2/internal/handlers"
	"Mohamed-Abdelrazeq/o-auth-2/internal/loaders"
	"Mohamed-Abdelrazeq/o-auth-2/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorMap struct {
	Message string `json:"message"`
}

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
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				ErrorMap{Message: "invalid credentials"},
			)
			return
		}
		ctx.JSON(200, user)
	})

	r.POST("/register", func(ctx *gin.Context) {
		user, err := authService.CreateUser(&database.CreateUserParams{Email: "mohamed@gmail.com", Password: "123456s"})
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				ErrorMap{Message: "email is used before"},
			)
			return
		}
		ctx.JSON(200, user)
	})

	// Run Server
	r.Run()
}
