package services

import (
	"Mohamed-Abdelrazeq/o-auth-2/internal/database"
	"Mohamed-Abdelrazeq/o-auth-2/internal/loaders"
	"context"
	"fmt"
)

// Identify Behaviour & Abstraction
type AuthService interface {
	GetUser(email string) (database.User, error)
	CreateUser(user database.CreateUserParams) (database.User, error)
}

// Abstraction
type AuthServiceInstance struct {
	DB *loaders.DatabaseInstance
}

// Initiator
func NewAuthSericeInstance(db *loaders.DatabaseInstance) AuthServiceInstance {
	return AuthServiceInstance{db}
}

// Behaviour
func (authServiceInstance AuthServiceInstance) GetUser(email string) (database.User, error) {
	user, err := authServiceInstance.DB.DB.GetUserByEmail(context.Background(), email)
	if err != nil {
		fmt.Println(err)
	}
	return user, err
}

func (authServiceInstance AuthServiceInstance) CreateUser(createUserParams *database.CreateUserParams) (database.User, error) {
	user, err := authServiceInstance.DB.DB.CreateUser(context.Background(), *createUserParams)
	if err != nil {
		fmt.Println(err)
	}
	return user, err
}
