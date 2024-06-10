package routes

import (
	"user-service/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CreateApp() *gin.Engine {
	app := gin.Default()
	// swagger endpoints
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// users endpoints
	app.POST("/users", func(c *gin.Context) {
		services.CreateUser(c)
	})

	app.GET("/users", func(c *gin.Context) {
		id := c.Query("id")
		username := c.Query("username")
		email := c.Query("email")
		services.GetUser(c, id, username, email)
	})

	return app
}
