package models

import "time"

type Posts struct {
	ID        uint      `json:"-" gorm:"primarykey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}

type CreatePostResponse struct {
	Status string `json:"status"`
	Post   Posts  `json:"post"`
}
