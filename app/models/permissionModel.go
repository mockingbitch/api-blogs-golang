package models

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	Id              int    `json:"id" gorm:"primaryKey"`
	Description     string `json:"description"`
	RolePermissions []RolePermission
	CreatedAt       time.Time      `json:"created_at" gorm:"<-:create"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}
