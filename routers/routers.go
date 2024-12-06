package routers

import (
	"net/http"

	"github.com/KarasunoAs9/RestAPI-Project/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(server *gin.Engine) {

	server.GET("/", func(ctx *gin.Context) {ctx.JSON(http.StatusOK, gin.H{"message": "default page"})})

	server.POST("/event/create-event", middleware.Authenticate, CreateEvent)
}