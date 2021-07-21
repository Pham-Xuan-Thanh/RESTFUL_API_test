package main

import (
	"golang-example/configs"
	"golang-example/entities"
	"golang-example/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// init app
	godotenv.Load()

	// connection to db mysql
	db, err := utils.ConnectToDatabaseMysql()

	if err != nil {
		defer utils.CloseConnectToDatabaseMysql(db)
	}
	gin.ForceConsoleColor()

	server := configs.InitServer(db)
	// Miragte db
	db.AutoMigrate(entities.Device{})
	err = server.Start()
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
