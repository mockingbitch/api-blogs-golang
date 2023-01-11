package models

import "gorm.io/gorm"

type RolePermission struct {
	gorm.Model
	Id           int  `json:"id" gorm:"primaryKey"`
	RoleId       int  `json:"role_id"`
	PermissionId int  `json:"permission_id"`
	Role         Role `gorm:"foreignKey:RoleId"`
	Permission   Role `gorm:"foreignKey:PermissionId"`
}
