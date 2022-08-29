package main

import (
	"github.com/EarnestScott/playground/users/controllers"
	"github.com/EarnestScott/playground/users/database"
	"github.com/EarnestScott/playground/users/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect("host=localhost user=postgresUser password=postgresPW dbname=postgresDB port=5455 sslmode=disable TimeZone=America/Los_Angeles")
	database.Migrate()
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
