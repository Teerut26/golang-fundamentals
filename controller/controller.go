package controller

import "github.com/gin-gonic/gin"

func HomePage(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}

var count = 0

func CountPage(context *gin.Context) {
	context.JSON(200, gin.H{
		"count": count,
	})
	count += 1
}
