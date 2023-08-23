package controller

import (
	"mytodoapp/database"
	"mytodoapp/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	database.DB.Find(&users)
	c.JSON(200, &users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	database.DB.Create(&user)
	c.JSON(200, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	database.DB.Where("id= ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	database.DB.Where("id=?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	database.DB.Save(&user)
	c.JSON(200, &user)

}
