package main

import (
	"testback/config"
	"testback/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.LoadEnv()

	database.InitDB()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	router.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
