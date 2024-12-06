package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(server *gin.Engine) {

	server.GET("/", func(ctx *gin.Context) {ctx.JSON(http.StatusOK, gin.H{"message": "default page"})})
}