package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Id              int    `json:"id" gorm:"primaryKey"`
	Description     string `json:"description"`
	RolePermissions []RolePermission
}
