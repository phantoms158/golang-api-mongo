package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.Default()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router) //add this

	router.Run(":" + port)
}
