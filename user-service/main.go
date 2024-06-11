package main

import (
	_ "user-service/docs"
	"user-service/routes"
)

// @title User Service API
// @version 1.0
// @description This service manages user accounts and profiles.

// @host localhost:8081
// @BasePath /
func main() {
	app := routes.CreateApp()

	app.Run()
}
