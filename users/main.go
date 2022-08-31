package main

import (
	"os"

	"github.com/EarnestScott/playground/users/controllers"
	"github.com/EarnestScott/playground/users/database"
	"github.com/EarnestScott/playground/users/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	conn := os.Getenv("DB_CONN")
	port := ":" + os.Getenv("PORT")

	database.Connect(conn)
	// database.Connect("host=localhost user=postgresUser password=postgresPW dbname=postgresDB port=5455 sslmode=disable TimeZone=UTC")
	database.Migrate()
	router := initRouter()
	router.Run(port)
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
