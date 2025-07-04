package database

import (
	"fmt"
	"log"
	"testback/config"
	"testback/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbUser := config.GetEnv("DB_USER", "root")
	dbPass := config.GetEnv("DB_PASS", "")
	dbHost := config.GetEnv("DB_HOST", "localhost")
	dbPort := config.GetEnv("DB_PORT", "3306")
	dbName := config.GetEnv("DB_NAME", "")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed Connect to Database")
	}
	fmt.Println("Database Sucessces Connected")

	err = DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Failed Auto Migrated Database", err)
	}

	fmt.Println("Auto Migrated Success")

}
