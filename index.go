package main

import (
	"fmt"
	"go-fundamentals/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	port := "91"

	router.GET("/home", controller.HomePage)
	router.GET("/count", controller.CountPage)

	fmt.Println("run on port : ", port)
	router.Run(":" + port)
}
