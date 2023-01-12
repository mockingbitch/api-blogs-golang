package models

import (
	"time"

	"gorm.io/gorm"
)

type RolePermission struct {
	Id           int            `json:"id" gorm:"primaryKey"`
	RoleId       int            `json:"role_id"`
	PermissionId int            `json:"permission_id"`
	Role         Role           `gorm:"foreignKey:RoleId"`
	Permission   Role           `gorm:"foreignKey:PermissionId"`
	CreatedAt    time.Time      `json:"created_at" gorm:"<-:create"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
