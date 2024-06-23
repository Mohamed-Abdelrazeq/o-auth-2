package helpers

import (
	"Mohamed-Abdelrazeq/o-auth-2/internal/models"
	"os"

	"github.com/golang-jwt/jwt"
)

// Generate New JWT
func NewAccessToken(claims models.UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

// Parse JWT
func ParseAccessToken(accessToken string) *models.UserClaims {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	return parsedAccessToken.Claims.(*models.UserClaims)
}
