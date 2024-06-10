package main

import (
	"os"
	_ "user-service/docs"
	routes "user-service/routes"
	"user-service/util"
)

// @title User Service API
// @version 1.0
// @description This service manages user accounts and profiles.

// @host localhost:8081
// @BasePath /
func main() {
	app := routes.CreateApp()
	if os.Getenv("MIGRATE") != "" {
		util.MigrateModels()
	}
	app.Run(":8081")
}
