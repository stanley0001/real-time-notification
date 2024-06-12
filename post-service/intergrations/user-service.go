package util

import (
	"net/http"
	"os"
	"post-service/util"

	"github.com/gin-gonic/gin"
)

func init() {
	util.LoadEnv()
}

var host = os.Getenv("USER_SERVICE_HOST")

func GetUserInfo(c *gin.Context) {
	var userId = c.Query("userId")
	//get user info from the gin context
	request, err := http.Get(host + "?userId=" + userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, request.Body)
}

func GetAllActiveUsers(c *gin.Context) {
	var status = c.Query("status")
	//get all active users from the gin context
	request, err := http.Get(host + "?status=" + status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, request.Body)
}
