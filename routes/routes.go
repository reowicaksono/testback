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
	router.PUT("/api/users/:id", middlewares.AuthMiddleware(), controllers.UpdateUser)
	router.DELETE("/api/users/:id", middlewares.AuthMiddleware(), controllers.DeleteUser)
	router.GET("/api/users", middlewares.AuthMiddleware(), controllers.FindUser)
	router.GET("/api/users/:id", middlewares.AuthMiddleware(), controllers.FindUserById)

	return router
}
