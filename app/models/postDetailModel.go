package models

import (
	"time"

	"gorm.io/gorm"
)

type PostDetail struct {
	Id        int            `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	PostId    int            `json:"post_id" gorm:"column:post_id"`
	Body      string         `json:"body"`
	Thumbnail *string        `json:"thumbnail"`
	CreatedAt time.Time      `json:"created_at" gorm:"<-:create"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
