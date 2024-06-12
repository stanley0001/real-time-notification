package main

import (
	_ "notification-service/docs"
	"notification-service/util"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Notification Service API
// @version 1.0
// @description This service generates and delivers real-time notifications to users.

// @host localhost:8082
// @BasePath /
func init() {
	util.LoadEnv()
	go util.ListenForEvents()
}

func main() {
	app := gin.Default()

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//users endpoints
	// app.POST("/users", services.CreateUser)
	// app.GET("/users/:id", services.GetUser)

	app.Run()
}
