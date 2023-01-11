package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Id         int  `json:"id" gorm:"primaryKey"`
	UserId     int  `gorm:"column:user_id"`
	Like       *int `json:"like;default:0"`
	PostDetail PostDetail
	User       User
}
