package util

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DbInstance *gorm.DB
var dbOnce sync.Once

func init() {
	LoadEnv()
}
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func ConnectDB() (*gorm.DB, error) {
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
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)

	DbInstance, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if os.Getenv("LOG_MODE") == "Info" {
		DbInstance.Logger.LogMode(logger.Info)
	}
	if err != nil {
		return nil, err
	}

	return DbInstance, nil
}

func GetDbConnection() *gorm.DB {
	dbOnce.Do(func() {
		var err error
		DbInstance, err = ConnectDB()
		if err != nil {
			log.Fatalf("Unable to connect to database: %v\n", err)
		}
	})

	return DbInstance
}
