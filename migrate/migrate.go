package main

import (
	initializers "GoginOrmDocker/Initializers"
	"GoginOrmDocker/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
