package services

import (
	"net/http"
	"post-service/models"
	"post-service/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var conn *gorm.DB = util.GetDbConnection()

// createPost godoc
// @Summary Create a post
// @Description Create a new post
// @Tags posts
// @Accept json
// @Produce json
// @Param post body models.Posts true "Post Data"
// @Success 201 {object} models.CreatePostResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /create-post [post]
func CreatePost(c *gin.Context) {
	var post models.Posts
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := conn.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	//TODO:fire notification event
	c.JSON(http.StatusCreated, post)
}

// UpdatePost godoc
// @Summary Update a post
// @Description update a post
// @Tags posts
// @Accept json
// @Produce json
// @Param post body models.Posts true "Post Data"
// @Success 201 {object} models.CreatePostResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /update-post [put]
func UpdatePost(c *gin.Context) {
	var post models.Posts
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := conn.First(&post, 1)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

//TODO: all other methods to be implemented here
