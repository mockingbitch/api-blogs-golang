package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"api-jwt-blogs/app/models"
	"api-jwt-blogs/database"
)

type Server struct {
	DB *gorm.DB
	// Router *mux.Router
}

func CreateRoles(ctx *gin.Context) {

}

func GetRoles(ctx *gin.Context) {
	role := models.Role{}
	// var server Server

	roles, err := role.FindAllRoles(database.Instance)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"roles": roles})
}
