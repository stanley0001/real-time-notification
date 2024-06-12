package models

import (
	"time"
)

type Users struct {
	ID        uint      `json:"-" gorm:"primarykey"`
	Name      string    `json:"name"`
	UserName  string    `json:"username" gorm:"uniqueIndex"`
	Email     string    `json:"email" gorm:"uniqueIndex"`
	Phone     int16     `json:"phone" gorm:"uniqueIndex"`
	Status    string    `json:"status"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}

type CreateUserResponse struct {
	Status string `json:"status"`
	User   Users  `json:"user"`
}
