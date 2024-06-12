package routes

import (
	"net/http"
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
	app.GET("/posts", func(c *gin.Context) {
		//get pagination params from the  gin context
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "paginated data to be returned here"})
	})
	//other endpoints, user posts,new posts...
	app.POST("/posts/create", func(c *gin.Context) {
		services.CreatePost(c)
	})
	app.PUT("/posts/update", func(c *gin.Context) {
		services.UpdatePost(c)
	})
	// comments endpoints
	app.GET("/comments", func(c *gin.Context) {
		//get pagination params from the  gin context
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "paginated data to be returned here"})
	})
	app.POST("/comments", func(c *gin.Context) {
		services.CreateComment(c)
	})
	//other endpoints, user comments,post comments new comments...

	return app
}
