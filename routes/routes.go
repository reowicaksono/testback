package routes

import (
	"testback/controllers"
	"testback/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// auth
	router.POST("/api/register", controllers.Register)

	router.POST("/api/login", controllers.Login)

	// users

	router.POST("/api/users", middlewares.AuthMiddleware(), controllers.CreateUser)

	router.GET("/api/users", middlewares.AuthMiddleware(), controllers.FindUser)

	return router
}
