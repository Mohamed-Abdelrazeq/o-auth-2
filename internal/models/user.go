package models

import "Mohamed-Abdelrazeq/o-auth-2/internal/database"

type CreateUserParams struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserParams struct {
	Email string `json:"email" binding:"required"`
}

func (data CreateUserParams) ConvertToDatabaseModel() *database.CreateUserParams {
	return &database.CreateUserParams{
		Email:    data.Email,
		Password: data.Password,
	}
}
