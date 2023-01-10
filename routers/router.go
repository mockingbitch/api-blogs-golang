package routers

import (
	"github.com/gin-gonic/gin"

	"api-jwt-blogs/app/http/controllers"
	"api-jwt-blogs/middlewares"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/login", controllers.Login)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
