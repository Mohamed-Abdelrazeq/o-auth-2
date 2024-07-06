package helpers

import (
	"Mohamed-Abdelrazeq/o-auth-2/internal/models"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

// Generate New JWT
func NewAccessToken(claims models.UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

// Parse JWT
func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
