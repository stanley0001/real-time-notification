package models

import (
	"time"
)

type Messages struct {
	Id         uint   `json:"-" gorm:"primarykey"`
	Content    string `json:"content"`
	UserFromId uint
	UserToId   uint
	UserFrom   Users     `gorm:"foreignKey:UserFromId;references:ID" json:"-"`
	UserTo     Users     `gorm:"foreignKey:UserToId;references:ID" json:"-"`
	Status     string    `json:"status"`
	SentAt     time.Time `json:"-" gorm:"autoCreateTime"`
	ReceivedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}

type SendMessageResponse struct {
	Status string   `json:"status"`
	User   Messages `json:"message"`
}
