package main

import (
	_ "post-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Post Service API
// @version 1.0
// @description This service handles creation, storage, and retrieval of posts.

// @host localhost:8080
// @BasePath /

func main() {
	app := gin.Default()

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//users endpoints
	// app.POST("/users", services.CreateUser)
	// app.GET("/users/:id", services.GetUser)

	app.Run(":8080")
}
