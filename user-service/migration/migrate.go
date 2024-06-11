package main

import (
	"log"
	models "user-service/models"
	"user-service/util"
)

func init() {
	util.LoadEnv()
	util.ConnectDB()
}
func main() {
	db := util.GetDbConnection()

	err := db.AutoMigrate(
		&models.Users{},
		&models.Following{},
		&models.Messages{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v\n", err)
	}
}
