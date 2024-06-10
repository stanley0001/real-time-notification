package util

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	models "user-service/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var dbInstance *gorm.DB
var dbOnce sync.Once

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func connectDB() (*gorm.DB, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetDbConnection() *gorm.DB {
	loadEnv()

	dbOnce.Do(func() {
		var err error
		dbInstance, err = connectDB()
		if err != nil {
			log.Fatalf("Unable to connect to database: %v\n", err)
		}
	})

	return dbInstance
}

func MigrateModels() {
	db := GetDbConnection()

	err := db.AutoMigrate(
		&models.User{},
		&models.Following{},
		&models.Messages{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v\n", err)
	}
}
