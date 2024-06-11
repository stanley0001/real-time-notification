package routes

import (
	"post-service/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CreateApp() *gin.Engine {
	app := gin.Default()
	// swagger endpoints
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// posts endpoints
	app.POST("/create-post", func(c *gin.Context) {
		services.CreatePost(c)
	})
	app.PUT("/update-post", func(c *gin.Context) {
		services.CreatePost(c)
	})

	// comments endpoints
	app.POST("/comments", func(c *gin.Context) {
		services.CreateComment(c)
	})

	return app
}
