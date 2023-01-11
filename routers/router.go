package routers

import (
	"github.com/gin-gonic/gin"

	"api-jwt-blogs/app/http/controllers"
	"api-jwt-blogs/middlewares"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	r := router.Group("/api/v1")
	{
		r.POST("/login", controllers.Login)
		r.POST("/register", controllers.RegisterUser)
		api := r.Group("/user").Use(middlewares.Auth())
		{
			api.GET("/ping", controllers.Ping)
		}
		admin := r.Group("/admin").Use(middlewares.Auth())
		{
			admin.GET("/roles", controllers.GetRoles)
		}

	}
	return router
}
