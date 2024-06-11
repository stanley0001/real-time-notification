package models

import "time"

type Comments struct {
	ID        uint      `json:"-" gorm:"primarykey"`
	Content   string    `json:"content"`
	PostID    uint      `json:"post_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}
