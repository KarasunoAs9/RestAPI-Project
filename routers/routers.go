package routers

import (
	"github.com/KarasunoAs9/RestAPI-Project/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(server *gin.Engine) {

	server.GET("/all-events", middleware.Authenticate, getAllEvents)
	server.POST("/event/create-event", middleware.Authenticate, CreateEvent)
	server.POST("/create-user", CreateUser)
	server.POST("/login", loginUser)
	server.PUT("/event/update/:id", middleware.Authenticate, updateEvent)
	server.DELETE("/event/delete/:id", middleware.Authenticate, deleteEvent)

}
