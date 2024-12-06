package routers

import (
	"log"
	"net/http"
	"strconv"

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

func updateEvent(ctx * gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 64)

	var event *models.Event

	err := ctx.ShouldBindBodyWithJSON(&event)

	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error with update event"})
		return
	}


	userId := ctx.GetInt64("userId")


	err = event.UpdateEvent(id, userId)

	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error with updating event"})
		return
	}


	ctx.JSON(http.StatusNoContent, gin.H{"message": "event was update sucsessfully"})

}

func deleteEvent(ctx *gin.Context) {
	var event *models.Event
	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 64)

	userId := ctx.GetInt64("userId")

	err := event.DeleteEvent(id, userId)

	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error with deliting event"})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message": "event was delete sucsessfully"})

}

func getAllEvents(ctx *gin.Context) {

}