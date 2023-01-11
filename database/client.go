package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"api-jwt-blogs/app/models"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}
func Migrate() {
	Instance.AutoMigrate(&models.Permission{})
	Instance.AutoMigrate(&models.Role{})
	Instance.AutoMigrate(&models.RolePermission{})
	Instance.AutoMigrate(&models.User{})
	Instance.AutoMigrate(&models.Post{})
	Instance.AutoMigrate(&models.PostDetail{})
	log.Println("Database Migration Completed!")
}
