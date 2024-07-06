package middlewares

import (
	"Mohamed-Abdelrazeq/o-auth-2/internal/helpers"
	"Mohamed-Abdelrazeq/o-auth-2/internal/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Check if request have authentication header
		authorization := ctx.GetHeader("Authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			ctx.JSON(
				http.StatusBadRequest,
				models.ErrorMap{Message: "no authentication header"},
			)
			return
		}
		// Check if request have a jwt token
		splits := strings.Split(authorization, " ")
		if len(splits) != 2 {
			ctx.JSON(
				http.StatusBadRequest,
				models.ErrorMap{Message: "no jwt token"},
			)
			return
		}
		// Check if request have valid token
		userClaims, err := helpers.VerifyToken(splits[1])
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				models.ErrorMap{Message: err.Error()},
			)
			return
		}

		ctx.Set("user_claims", userClaims)
		ctx.Next()
	}
}
