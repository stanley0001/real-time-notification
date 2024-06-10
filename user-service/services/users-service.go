package services

import (
	"net/http"
	_ "user-service/docs"
	models "user-service/models"

	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

// createUser godoc
// @Summary Create a user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User Data"
// @Success 201 {object} models.CreateUserResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /users [post]
func CreateUser(c *gin.Context, conn *pgx.Conn) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := conn.Exec(context.Background(), "INSERT INTO users (id,name, email,username) VALUES ($1, $2, $3, $4)", user.ID, user.Name, user.Email, user.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
// @Success 200 {object} models.User
// @Failure 404 {object} models.ErrorResponse
// @Router /users [get]
func GetUser(c *gin.Context, conn *pgx.Conn, id string, userName string, email string) {
	// id := c.Param("id")
	// email := c.Param("email")
	// userName := c.Param("username")

	var user, err = getUser(id, email, userName, conn)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, user)
}

func getUser(id string, email string, userName string, conn *pgx.Conn) (models.User, error) {
	var user models.User
	if id != "" {
		err := conn.QueryRow(context.Background(), "SELECT id, name, email,username FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email, &user.UserName)
		return user, err
	}
	if userName != "" {
		err := conn.QueryRow(context.Background(), "SELECT id, name, email,username FROM users WHERE username=$1", userName).Scan(&user.ID, &user.Name, &user.Email, &user.UserName)
		return user, err
	}
	if email != "" {
		err := conn.QueryRow(context.Background(), "SELECT id, name, email,username FROM users WHERE email=$1", email).Scan(&user.ID, &user.Name, &user.Email, &user.UserName)
		return user, err
	}
	return user, nil
}
