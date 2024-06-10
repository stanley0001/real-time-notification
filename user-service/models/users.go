package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string    `json:"name" gorm:"tableName:users"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Phone     uint      `json:"phone"`
	Status    string    `json:"status"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateUserResponse struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}
