package main

import (
	apps "Mohamed-Abdelrazeq/o-auth-2/internal/application"
	"Mohamed-Abdelrazeq/o-auth-2/internal/handlers"
	"Mohamed-Abdelrazeq/o-auth-2/internal/loaders"
	"Mohamed-Abdelrazeq/o-auth-2/internal/models"
	"Mohamed-Abdelrazeq/o-auth-2/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	// Load Env & DB
	loaders.LoadEnv()
	db := loaders.LoadDB()

	// Init AuthService
	authService := services.NewAuthSericeInstance(db)
	authApp := apps.NewAuthApplicationInstance(authService)

	// Create Router
	r := gin.Default()

	// Test Router
	r.GET("/ping", handlers.Ping)

	// Auth Router
	r.POST("/login", func(ctx *gin.Context) {
		// Binding
		var loginUserParams models.LoginUserParams
		ctx.Bind(&loginUserParams)
		// Validation
		validate := validator.New()
		if err := validate.Struct(loginUserParams); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				models.ErrorMap{Message: "invalid body"},
			)
			return
		}
		token, err := authApp.Login(loginUserParams)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				models.ErrorMap{Message: "invalid credintials"},
			)
			return
		}

		// Return
		ctx.JSON(200, token)
	})

	r.POST("/register", func(ctx *gin.Context) {
		// Binding
		var createUserParams models.CreateUserParams
		ctx.Bind(&createUserParams)
		// Validation
		validate := validator.New()
		if err := validate.Struct(createUserParams); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				models.ErrorMap{Message: "invalid body"},
			)
			return
		}
		//
		token, err := authApp.Register(createUserParams)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				models.ErrorMap{Message: "invalid credintials"},
			)
			return
		}

		// Return
		ctx.JSON(200, token)
	})

	// Run Server
	r.Run()
}
