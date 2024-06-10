package models

import (
	"gorm.io/gorm"
)

type Following struct {
	gorm.Model
	FollowedID uint
	FollowerID uint
	Followed   User `gorm:"foreignkey:FollowedID"`
	Follower   User `gorm:"foreignkey:FollowerID"`
}
