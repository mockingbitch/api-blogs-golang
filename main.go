package main

import (
	"api-jwt-blogs/config"
	"api-jwt-blogs/database"
	"api-jwt-blogs/routers"
)

func main() {
	// Initialize Database
	dsn := config.SetupDatabaseConnection()
	database.Connect(dsn)
	database.Migrate()

	// Initialize Router
	router := routers.InitRouter()
	router.Run(":9090")
}
