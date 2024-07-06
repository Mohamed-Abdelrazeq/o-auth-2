package models

import (
	"Mohamed-Abdelrazeq/o-auth-2/internal/database"

	_ "github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

type CreateUserParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=128"`
}

type LoginUserParams struct {
	CreateUserParams
}

type UserClaims struct {
	Id        int   `json:"id" binding:"required"`
	IsActive  bool  `json:"is_active" binding:"required"`
	ExpiresAt int64 `json:"expires_at" binding:"required"`
	jwt.StandardClaims
}

type Token struct {
	Token string `json:"token"`
}

func (data CreateUserParams) ConvertToDatabaseModel() *database.CreateUserParams {
	return &database.CreateUserParams{
		Email:    data.Email,
		Password: data.Password,
	}
}
