package services

import (
	"net/http"
	models "user-service/models"

	"github.com/gin-gonic/gin"
)

// sendMessage godoc
// @Summary Send a message
// @Description Send a new message
// @Tags messages
// @Accept json
// @Produce json
// @Param message body models.Messages true "Message"
// @Success 201 {object} models.SendMessageResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /messages [post]
func SendMessage(c *gin.Context) {
	var message models.Messages

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	result := conn.Create(&message)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	//fire notification event
	c.JSON(http.StatusCreated, message)
}

func userOutBox(user models.Users) []models.Messages {
	var userOutBox []models.Messages
	userOutBox = append(userOutBox, models.Messages{UserToId: user.ID}) //TODO :logic to be implemed
	//get outbox from messages table using user id
	return userOutBox
}
func userInBox(user models.Users) []models.Messages {
	var userIntBox []models.Messages
	userIntBox = append(userIntBox, models.Messages{UserToId: user.ID}) //TODO :logic to be implemed
	//get outbox from messages table using user id
	return userIntBox
}
