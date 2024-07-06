package helpers

import (
	"Mohamed-Abdelrazeq/o-auth-2/internal/models"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// Generate New JWT
func NewAccessToken(claims models.UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

// Parse JWT
func VerifyToken(tokenString string) (*models.UserClaims, error) {
	var userClaims = &models.UserClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if err != nil {
		return userClaims, err
	}

	if !token.Valid {
		return userClaims, errors.New("invalid token")
	}

	userClaims = token.Claims.(*models.UserClaims)

	if userClaims.ExpiresAt != 0 && userClaims.ExpiresAt < time.Now().Unix() {
		return userClaims, errors.New("token expired")
	}

	return userClaims, nil
}
