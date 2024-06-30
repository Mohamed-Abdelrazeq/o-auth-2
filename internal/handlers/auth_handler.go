package handlers

import (
	apps "Mohamed-Abdelrazeq/o-auth-2/internal/application"
	"Mohamed-Abdelrazeq/o-auth-2/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type AuthHandlerInstance struct {
	authApp apps.AuthApplication
}

func NewAuthHandlerInstance(authApp apps.AuthApplication) AuthHandler {
	return AuthHandlerInstance{authApp: authApp}
}

// Register With Email & Password
func (authHandler AuthHandlerInstance) Register(ctx *gin.Context) {
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
	token, err := authHandler.authApp.Register(createUserParams)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			models.ErrorMap{Message: "invalid credintials"},
		)
		return
	}

	// Return
	ctx.JSON(200, token)
}

// Login With Email & Password
func (authHandler AuthHandlerInstance) Login(ctx *gin.Context) {

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
	token, err := authHandler.authApp.Login(loginUserParams)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			models.ErrorMap{Message: "invalid credintials"},
		)
		return
	}

	// Return
	ctx.JSON(200, token)
}
