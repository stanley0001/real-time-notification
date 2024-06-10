package main

import (
	"time"
	_ "user-service/docs"
	services "user-service/services"

	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title User Service API
// @version 1.0
// @description This service manages user accounts and profiles.

// @host localhost:8081
// @BasePath /

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func connectDB() (*pgx.Conn, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func main() {
	loadEnv()

	conn, err := connectDB()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	app := gin.Default()

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//users endpoints
	app.POST("/users", func(c *gin.Context) {
		services.CreateUser(c, conn)
	})

	app.GET("/users", func(c *gin.Context) {
		id := c.Query("id")
		username := c.Query("username")
		email := c.Query("email")
		services.GetUser(c, conn, id, username, email)
		// services.GetUser(c, conn)
	})

	app.Run(":8081")
}
