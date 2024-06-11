package main

import (
	"log"
	models "post-service/models"
	"post-service/util"
)

func init() {
	util.LoadEnv()
	util.ConnectDB()
}
func main() {
	db := util.GetDbConnection()

	err := db.AutoMigrate(
		&models.Posts{},
		&models.Comments{},
		&models.Reactions{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v\n", err)
	}
}
