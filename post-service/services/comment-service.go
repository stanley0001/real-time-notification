package services

import (
	"net/http"
	"post-service/models"

	"github.com/gin-gonic/gin"
)

// createComment godoc
// @Summary Create a comment
// @Description Create a new comment
// @Tags comments
// @Accept json
// @Produce json
// @Param comment body models.Comments true "Comment Data"
// @Success 201 {object} models.CreateCommentResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /comments [post]
func CreateComment(c *gin.Context) {
	var comment models.Comments
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := conn.Create(&comment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	//TODO:fire notification event
	c.JSON(http.StatusCreated, comment)
}
