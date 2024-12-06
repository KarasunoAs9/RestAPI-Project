package routers

import (
	"github.com/KarasunoAs9/RestAPI-Project/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(server *gin.Engine) {

	server.GET("/all-events", middleware.Authenticate, GetAllEvents)
	server.POST("/event/create-event", middleware.Authenticate, CreateEvent)
	server.POST("/create-user", CreateUser)
	server.POST("/login", loginUser)
	server.PUT("/update-event/:id", middleware.Authenticate, UpdateEvent)
	server.DELETE("/delete-event/:id", middleware.Authenticate, UpdateEvent)

}
