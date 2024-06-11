package main

import (
	_ "post-service/docs"
	"post-service/routes"
)

// @title Post Service API
// @version 1.0
// @description This service handles creation, storage, and retrieval of posts.

// @host localhost:8080
// @BasePath /

func main() {
	app := routes.CreateApp()

	app.Run()
}
