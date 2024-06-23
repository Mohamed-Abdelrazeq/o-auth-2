package models

import (
	"Mohamed-Abdelrazeq/o-auth-2/internal/database"

	"github.com/golang-jwt/jwt"
)

type CreateUserParams struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserParams struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserClaims struct {
	Id int `json:"id" binding:"required"`
	jwt.StandardClaims
}

type Token struct {
	Token string `json:"toke"`
}

func (data CreateUserParams) ConvertToDatabaseModel() *database.CreateUserParams {
	return &database.CreateUserParams{
		Email:    data.Email,
		Password: data.Password,
	}
}
