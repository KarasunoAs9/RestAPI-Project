package routers

import (
	"log"
	"net/http"

	"github.com/KarasunoAs9/RestAPI-Project/models"
	"github.com/KarasunoAs9/RestAPI-Project/utils"
	"github.com/gin-gonic/gin"
)

var user *models.User

func CreateUser(ctx *gin.Context) {
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error with creating user"})
		return
	}
	token, err := utils.GenerateToken(user.Username, user.ID)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error with generate token"})
		return
	}

	user.Save()
	ctx.JSON(http.StatusCreated, gin.H{"message": "user created sucsessfully", "token": token})
}