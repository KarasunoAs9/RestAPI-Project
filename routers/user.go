package routers

import (
	"log"
	"net/http"

	"github.com/KarasunoAs9/RestAPI-Project/models"
	"github.com/KarasunoAs9/RestAPI-Project/utils"
	"github.com/gin-gonic/gin"
)



func CreateUser(ctx *gin.Context) {
	var user *models.User
	
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error with creating user"})
		return
	}

	user.Save()
	ctx.JSON(http.StatusCreated, gin.H{"message": "user created sucsessfully"})
}

func loginUser(ctx *gin.Context) {
	var user *models.User

	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error with login"})
		return
	}

	err := user.VerifyUser()

	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "incorrect password or login"})
		return
	}

	token, err := utils.GenerateToken(user.Username, user.ID)

	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error with generate token"})
	}


	ctx.JSON(http.StatusOK, gin.H{"message": "login sucsessful", "token": token})

}

