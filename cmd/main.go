package main

import (
	"Mohamed-Abdelrazeq/o-auth-2/internal/handlers"
	"Mohamed-Abdelrazeq/o-auth-2/internal/loaders"
	"Mohamed-Abdelrazeq/o-auth-2/internal/models"
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
		// Validation
		var loginUserParams models.LoginUserParams
		if err := ctx.ShouldBind(&loginUserParams); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				ErrorMap{Message: err.Error()},
			)
			return
		}
		// Logic
		dbUser, err := authService.GetUser(loginUserParams.Email)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				ErrorMap{Message: "invalid credentials"},
			)
			return
		}
		// Return
		ctx.JSON(200, dbUser)
	})

	r.POST("/register", func(ctx *gin.Context) {
		// Validation
		var createUserParams models.CreateUserParams
		if err := ctx.ShouldBind(&createUserParams); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				ErrorMap{Message: err.Error()},
			)
			return
		}
		// Logic
		dbUser, err := authService.CreateUser(createUserParams.ConvertToDatabaseModel())
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				ErrorMap{Message: err.Error()},
			)
			return
		}
		// Return
		ctx.JSON(200, dbUser)
	})

	// Run Server
	r.Run()
}
