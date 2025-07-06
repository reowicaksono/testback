package main

import (
	"testback/config"
	"testback/database"
	"testback/routes"
)

func main() {

	config.LoadEnv()

	database.InitDB()

	r := routes.SetupRouter()

	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
