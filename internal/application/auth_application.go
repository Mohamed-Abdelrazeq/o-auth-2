package apps

import (
	"Mohamed-Abdelrazeq/o-auth-2/internal/helpers"
	"Mohamed-Abdelrazeq/o-auth-2/internal/models"
	"Mohamed-Abdelrazeq/o-auth-2/internal/services"
	"errors"
	"time"
)

type AuthApplication interface {
	Login(loginUserParams models.LoginUserParams) (models.Token, error)
	Register(createUserParams models.CreateUserParams) (models.Token, error)
}

type AuthApplicationInstance struct {
	authService services.AuthService
}

func NewAuthApplicationInstance(authService services.AuthService) AuthApplication {
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
	token.Token, err = helpers.NewAccessToken(
		models.UserClaims{
			Id:        int(dbUser.ID),
			IsActive:  dbUser.IsActive,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	)
	if err != nil {
		return token, err
	}
	return token, err
}

func (authApplication AuthApplicationInstance) Register(createUserParams models.CreateUserParams) (models.Token, error) {
	var token models.Token
	// Hashing
	createUserParams.Password, _ = helpers.HashPassword(createUserParams.Password)
	// DB
	dbUser, err := authApplication.authService.CreateUser(createUserParams.ConvertToDatabaseModel())
	if err != nil {
		return token, err

	}
	// Token
	token.Token, err = helpers.NewAccessToken(
		models.UserClaims{
			Id:       int(dbUser.ID),
			IsActive: dbUser.IsActive,
		},
	)
	if err != nil {
		return token, err
	}
	return token, nil
}
