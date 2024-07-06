package handlers

import "github.com/gin-gonic/gin"

func AuthenticationTest(ctx *gin.Context) {
	userClaims, exists := ctx.Get("user_claims")
	if !exists {
		return
	}
	ctx.JSON(
		200,
		gin.H{
			"message":     "Authenticated User",
			"user_claims": userClaims,
			"exists":      exists,
		},
	)
}
