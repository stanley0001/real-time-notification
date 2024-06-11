package models

import (
	"time"
)

type Following struct {
	ID         uint `json:"-" gorm:"primarykey"`
	FollowedID uint
	FollowerID uint
	Followed   Users     `gorm:"foreignKey:FollowedID;references:ID" json:"-"`
	Follower   Users     `gorm:"foreignKey:FollowerID;references:ID" json:"-"`
	CreatedAt  time.Time `json:"-" gorm:"autoCreateTime"`
	UpdateAt   time.Time `json:"-" gorm:"autoUpdateTime"`
}
