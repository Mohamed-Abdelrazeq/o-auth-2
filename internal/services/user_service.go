package services

import (
	"Mohamed-Abdelrazeq/o-auth-2/internal/database"
	"Mohamed-Abdelrazeq/o-auth-2/internal/loaders"
)

// Identify Behaviour & Abstraction
type AuthService interface {
	GetUser(email string) (database.User, error)
	CreateUser(user database.User) error
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
	user := database.User{Email: email}
	// response := authServiceInstance.DB.DB.Where("email = ?", email).First(&user)
	return user, nil
}

func (authServiceInstance AuthServiceInstance) CreateUser(user *database.User) error {
	// response := authServiceInstance.DB.DB.Create(user)
	return nil
}
