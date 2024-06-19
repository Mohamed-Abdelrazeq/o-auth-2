package services

import (
	"Mohamed-Abdelrazeq/o-auth-2/loaders"
	"Mohamed-Abdelrazeq/o-auth-2/models"
)

// Identify Behaviour & Abstraction
type AuthService interface {
	GetUser(email string) (models.User, error)
	CreateUser(user models.User) error
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
func (authServiceInstance AuthServiceInstance) GetUser(email string) (models.User, error) {
	user := models.User{Email: email}
	response := authServiceInstance.DB.DB.Where("email = ?", email).First(&user)
	return user, response.Error
}

func (authServiceInstance AuthServiceInstance) CreateUser(user *models.User) error {
	response := authServiceInstance.DB.DB.Create(user)
	return response.Error
}
