package main

import (
	"github.com/KarasunoAs9/RestAPI-Project/db"
	"github.com/KarasunoAs9/RestAPI-Project/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	routers.RegisterRouters(server)
	db.InitDB()
	
	server.Run(":8080")
}