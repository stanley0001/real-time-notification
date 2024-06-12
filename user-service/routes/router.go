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

	//messaging endpoints
	app.POST("/messages", func(c *gin.Context) {
		services.SendMessage(c)
	})

	//following endpoints
	app.GET("/users/follow", func(c *gin.Context) {
		followerId := c.Query("follower")
		followedId := c.Query("id")
		services.FollowUser(c, followerId, followedId)
	})
	app.GET("/users/followers", func(c *gin.Context) {
		services.GetFollowers(c)
	})
	//auth endpoints
	app.POST("/users/auth", func(c *gin.Context) {
		services.Authenticate(c)
	})
	return app
}
