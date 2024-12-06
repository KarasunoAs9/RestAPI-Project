package routers

import (
	"log"
	"net/http"

	"github.com/KarasunoAs9/RestAPI-Project/models"
	"github.com/gin-gonic/gin"
)




func CreateEvent(ctx *gin.Context) {

	var event *models.Event

	err := ctx.ShouldBindBodyWithJSON(&event)

	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error with event"})
		return
	}

	userId := ctx.GetInt64("userId")

	event.UserID = userId

	if err = event.Save(); err != nil {
		log.Fatal(userId, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "event was created sucsessfully", "event": event})
}