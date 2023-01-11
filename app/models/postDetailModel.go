package models

import "gorm.io/gorm"

type PostDetail struct {
	gorm.Model
	Id        int     `json:"id" gorm:"primaryKey"`
	Title     string  `json:"title"`
	PostId    int     `json:"post_id" gorm:"column:post_id"`
	Body      string  `json:"body"`
	Thumbnail *string `json:"thumbnail"`
}
