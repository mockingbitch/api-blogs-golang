package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	Id              int    `json:"id" gorm:"primaryKey"`
	Name            string `json:"name"`
	RolePermissions []RolePermission
	CreatedAt       time.Time      `json:"created_at" gorm:"<-:create"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
	// Users           []User
}

func (r *Role) SaveRole(db *gorm.DB) (*Role, error) {

	var err error
	err = db.Debug().Create(&r).Error
	if err != nil {
		return &Role{}, err
	}
	return r, nil
}

func (r *Role) FindAllRoles(db *gorm.DB) (*[]Role, error) {
	var err error
	roles := []Role{}
	err = db.Debug().Model(&Role{}).Limit(100).Find(&roles).Error
	if err != nil {
		return &[]Role{}, err
	}
	return &roles, err
}

func (r *Role) FindRoleByID(db *gorm.DB, uid uint32) (*Role, error) {
	var err error
	err = db.Debug().Model(Role{}).Where("id = ?", uid).Take(&r).Error
	if err != nil {
		return &Role{}, err
	}
	// if gorm.IsRecordNotFoundError(err) {
	// 	return &Role{}, errors.New("Role Not Found")
	// }
	return r, err
}

func (r *Role) UpdateAUser(db *gorm.DB, uid uint32) (*Role, error) {
	db = db.Debug().Model(&Role{}).Where("id = ?", uid).Take(&Role{}).UpdateColumns(
		map[string]interface{}{
			"name":      r.Name,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Role{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&Role{}).Where("id = ?", uid).Take(&r).Error
	if err != nil {
		return &Role{}, err
	}
	return r, nil
}

func (r *Role) DeleteARole(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Role{}).Where("id = ?", uid).Take(&Role{}).Delete(&Role{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
