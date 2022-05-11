package main

import (
	"go-fundamentals/configs"
	"go-fundamentals/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	port := "91"
	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router)

	router.Run(":" + port)
}
