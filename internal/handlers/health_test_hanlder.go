package handlers

import "github.com/gin-gonic/gin"

func HealthTest(ctx *gin.Context) {
	ctx.JSON(
		200,
		gin.H{
			"message": "pong",
		},
	)
}
