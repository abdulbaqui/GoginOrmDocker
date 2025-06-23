package main

import (
	initializers "GoginOrmDocker/Initializers"
	"GoginOrmDocker/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()

	// Try to connect to database, but don't fail if it's not available
	err := initializers.ConnectToDB()
	if err != nil {
		log.Printf("Warning: Database connection failed: %v", err)
		log.Println("Application will start without database connection")
	}
}

func main() {
	r := gin.Default()
	r.POST("/user", controllers.PostCreate)
	r.GET("/user", controllers.PostIndex)
	r.GET("/user/:id", controllers.GetSpecific)
	r.DELETE("/user/:id", controllers.Delete)
	r.PUT("/user/:id", controllers.UpdateUser)
	r.Run()
}
