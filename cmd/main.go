package main

import (
	apps "Mohamed-Abdelrazeq/o-auth-2/internal/application"
	"Mohamed-Abdelrazeq/o-auth-2/internal/handlers"
	"Mohamed-Abdelrazeq/o-auth-2/internal/helpers"
	"Mohamed-Abdelrazeq/o-auth-2/internal/loaders"
	"Mohamed-Abdelrazeq/o-auth-2/internal/models"
	"Mohamed-Abdelrazeq/o-auth-2/internal/services"
	"net/http"
	"strings"

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
	r.Use(AuthenticateUser())
	r.GET("/protected", handlers.AuthenticationTest)

	// Run Server
	r.Run()
}

func AuthenticateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Check if request have authentication header
		authorization := ctx.GetHeader("Authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			ctx.JSON(
				http.StatusBadRequest,
				models.ErrorMap{Message: "no authentication header"},
			)
			return
		}
		// Check if request have a jwt token
		splits := strings.Split(authorization, " ")
		if len(splits) != 2 {
			ctx.JSON(
				http.StatusBadRequest,
				models.ErrorMap{Message: "no jwt token"},
			)
			return
		}
		// Check if request have valid token
		userClaims, err := helpers.VerifyToken(splits[1])
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				models.ErrorMap{Message: "invalid token"},
			)
			return
		}
		ctx.Set("user_claims", userClaims)
		ctx.Next()
	}
}
