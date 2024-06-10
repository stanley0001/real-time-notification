package models

import (
	"time"

	"gorm.io/gorm"
)

type Messages struct {
	gorm.Model
	Content    string `json:"content"`
	UserFromId uint
	UserToId   uint
	UserFrom   User      `gorm:"foreignkey:UserFromId" json:"from"`
	UserTo     User      `gorm:"foreignkey:UserToId" json:"to"`
	Status     string    `json:"status"`
	SentAt     time.Time `json:"sentAt"`
	ReceivedAt time.Time `json:"receivedAt"`
}
