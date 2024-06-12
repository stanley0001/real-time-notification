package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"post-service/models"
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

// Auth godoc
// @Summary user authentication
// @Description authenticate user
// @Tags auth
// @Produce json
// @Param auth body models.Authentication true "Auth Data"
// @Success 200 {object} []models.AuthenticationResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /users/auth [post]
func Authenticate(c *gin.Context) {
	var user models.Authentication
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//auth logic here
	jsonData, _ := json.Marshal(user)
	request, _ := http.NewRequest("POST", host+"/auth", bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")
	// request.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var authResponse models.AuthenticationResponse
	if err := json.Unmarshal(bodyBytes, &authResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Response from user auth: %v", authResponse)
	c.JSON(http.StatusOK, authResponse)
}
