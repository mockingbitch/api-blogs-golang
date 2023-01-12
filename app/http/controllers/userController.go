package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api-jwt-blogs/app/models"
	"api-jwt-blogs/database"
)

type User struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func GetUsers(ctx *gin.Context) {
	user := models.User{}
	users, err := user.FindAllUsers(database.Instance)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"roles": users})
}
