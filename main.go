package main

import (
	initializers "GoginOrmDocker/Initializers"
	"GoginOrmDocker/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
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
