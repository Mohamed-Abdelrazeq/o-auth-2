package apps

import (
	"Mohamed-Abdelrazeq/o-auth-2/internal/helpers"
	"Mohamed-Abdelrazeq/o-auth-2/internal/models"
	"Mohamed-Abdelrazeq/o-auth-2/internal/services"
	"errors"
)

type AuthApplication interface {
	Login(loginUserParams models.LoginUserParams) (models.Token, error)
	Register(createUserParams models.CreateUserParams) (models.Token, error)
}

type AuthApplicationInstance struct {
	authService *services.AuthServiceInstance
}

func NewAuthApplicationInstance(authService *services.AuthServiceInstance) AuthApplicationInstance {
	return AuthApplicationInstance{authService: authService}
}

func (authApplication AuthApplicationInstance) Login(loginUserParams models.LoginUserParams) (models.Token, error) {
	var token models.Token
	// DB
	dbUser, err := authApplication.authService.GetUser(loginUserParams.Email)
	if err != nil {
		return token, err
	}
	// Hashing
	isVerified := helpers.VerifyPassword(loginUserParams.Password, dbUser.Password)
	if !isVerified {
		return token, errors.New("invalid password")
	}
	// Token
	token.Token, err = helpers.NewAccessToken(models.UserClaims{Id: int(dbUser.ID)})
	if err != nil {
		return token, err
	}
	return token, err
}
