package main

import (
	"flag"
	"fmt"

	"github.com/EarnestScott/playground/users/controllers"
	"github.com/EarnestScott/playground/users/database"
	"github.com/EarnestScott/playground/users/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {

	conn := flag.String("conn", "", "Specify db connection string")

	// Declare a flag called age with default value of 0 and a help message
	port := flag.Int("port", 0, "Specify connection port")

	// Enable command-line parsing
	flag.Parse()

	if *conn == "" {
		panic("No connection string specified")
	}
	if *port == 0 {
		panic("No port number specified")
	}

	portStr := fmt.Sprint(":%d", *port)
	database.Connect(*conn)
	database.Migrate()
	router := initRouter()
	router.Run(portStr)
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
