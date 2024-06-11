package models

import "time"

type Reactions struct {
	ID        uint      `json:"-" gorm:"primarykey"`
	Type      string    `json:"type"`
	PostID    uint      `json:"post_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}
