package services

import (
	"errors"
	"net/http"
	_ "user-service/docs"
	models "user-service/models"
	"user-service/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var conn *gorm.DB = util.GetDbConnection()

// createUser godoc
// @Summary Create a user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.Users true "User Data"
// @Success 201 {object} models.CreateUserResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := conn.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.Users true "User Data"
// @Success 201 {object} models.CreateUserResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /users [put]
func UpdateUser(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// _, err := conn.Exec(context.Background(), "INSERT INTO users (id,name, email,username) VALUES ($1, $2, $3, $4)", user.ID, user.Name, user.Email, user.UserName)
	result := conn.First(&user, 1)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// getUser godoc
// @Summary Get a user
// @Description Get user details by id, username, or email
// @Tags users
// @Produce json
// @Param id query string false "User ID"
// @Param username query string false "User Username"
// @Param email query string false "User Email"
// @Success 200 {object} models.Users
// @Failure 404 {object} models.ErrorResponse
// @Router /users [get]
func GetUser(c *gin.Context, id string, userName string, email string) {
	// id := c.Param("id")
	// email := c.Param("email")
	// userName := c.Param("username")

	var user, err = getUser(id, email, userName)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, user)
}

// follow user godoc
// @Summary follow user
// @Description Create a folloing entry between two users
// @Tags following
// @Produce json
// @Param id query string false "User ID"
// @Param follower query string false "Follower ID"
// @Success 200 {object} models.Following
// @Failure 404 {object} models.ErrorResponse
// @Router /users/follow [get]
func FollowUser(c *gin.Context, followerId string, userId string) {
	var follower, user, err = fetchFollowedAndFollowers(followerId, userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return

	}
	var followingEntry = models.Following{FollowedID: user.ID, FollowerID: follower.ID}
	result := conn.Create(&followingEntry)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	//fire notification event
	c.JSON(http.StatusCreated, followingEntry)
}

// get followers godoc
// @Summary user followers
// @Description get all followers
// @Tags following
// @Produce json
// @Param id query string false "User ID"
// @Success 200 {object} []models.Following
// @Failure 404 {object} models.ErrorResponse
// @Router /users/followers [get]
func GetFollowers(c *gin.Context) {
	userId := c.Param("id")
	var user models.Users
	err := conn.First(&user, userId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	followers := fetchFollowers(user)
	c.JSON(http.StatusOK, followers)
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
	//auth logic here
	c.JSON(http.StatusOK, models.AuthenticationResponse{Status: "success", Token: "demo token"})
}

func fetchFollowedAndFollowers(userId string, followerId string) (models.Users, models.Users, error) {
	followers, err := getUser(followerId, "", "")
	if err != nil {
		return models.Users{}, models.Users{}, err
	}
	if len(followers) == 0 {
		return models.Users{}, models.Users{}, errors.New("follower not found")
	}
	follower := followers[0]

	// Fetch followed
	followeds, err := getUser(userId, "", "")
	if err != nil {
		return models.Users{}, models.Users{}, err
	}
	if len(followeds) == 0 {
		return models.Users{}, models.Users{}, errors.New("followed not found")
	}
	followed := followeds[0]

	return follower, followed, nil
}

func fetchFollowers(user models.Users) []models.Users {
	var followers []models.Users
	followers = append(followers, models.Users{ID: user.ID}) //TODO :logic and pagination to be implemed
	//get followers from followers table using user id
	return followers
}

func getUser(id string, email string, userName string) ([]models.Users, error) {
	var user models.Users
	var users []models.Users

	if id != "" {
		result := conn.First(&user, id)
		users = append(users, user)
		return users, result.Error
	}
	if userName != "" {
		result := conn.Where("user_name = ?", userName).Find(&users)
		return users, result.Error
	}
	if email != "" {
		result := conn.Where("Email = ?", email).Find(&users)
		return users, result.Error
	}
	return users, nil
}
