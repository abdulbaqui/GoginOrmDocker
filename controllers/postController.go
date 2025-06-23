package controllers

import (
	initializers "GoginOrmDocker/Initializers"
	"GoginOrmDocker/models"
	"log"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	// Get Data
	var body struct {
		Name   string
		Age    int
		Gender bool
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Create Post

	user := models.User{Name: body.Name, Age: body.Age, Gender: body.Gender}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		log.Fatal(result.Error)
		return
	}

	//Return it

	c.JSON(200, gin.H{
		"user": user,
	})
}

func PostIndex(c *gin.Context) {
	// Get all Objects

	var user []models.User

	initializers.DB.Find(&user)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func Delete(c *gin.Context) {
	// Get Data
	id := c.Param("id")
	var user models.User
	initializers.DB.Delete(&user, id)
}

func GetSpecific(c *gin.Context) {
	//Get Id of y URL
	id := c.Param("id")
	var user models.User

	initializers.DB.Find(&user, id)

	c.JSON(200, gin.H{
		"user": user,
	})

}

// Update User
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Name   string
		Age    int
		Gender bool
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	initializers.DB.First(&user, id)

	initializers.DB.Model(&user).Updates(models.User{
		Name:   body.Name,
		Age:    body.Age,
		Gender: body.Gender,
	})

	c.JSON(200, gin.H{
		"user": user,
	})

}
