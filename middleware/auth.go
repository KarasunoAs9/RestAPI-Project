package middleware

import (
	"net/http"

	"github.com/KarasunoAs9/RestAPI-Project/utils"
	"github.com/gin-gonic/gin"
)


func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "user not authenticated"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unavailable token"})
		return
	}

	ctx.Set("userId", userId)

	ctx.Next()
}