package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	Id         int  `json:"id" gorm:"primaryKey"`
	UserId     int  `gorm:"column:user_id"`
	Like       *int `json:"like;default:0"`
	PostDetail PostDetail
	User       User
	CreatedAt  time.Time      `json:"created_at" gorm:"<-:create"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
