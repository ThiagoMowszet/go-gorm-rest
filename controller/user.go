package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/thiagomowszet/go-gorm-rest/config"
    "github.com/thiagomowszet/go-gorm-rest/models"
)

func GetUsers(c *gin.Context) {
    users := []models.User{}
    config.DB.Find(&users)
    // c.String(200, "Hello World!")
    c.JSON(200, &users)
}

func CreateUser(c *gin.Context) {
    var user models.User
    c.BindJSON(&user)
    config.DB.Create(&user)
    c.JSON(200, &user)
}

func DeleteUser(c *gin.Context) {
    var user models.User
    config.DB.Where("id = ?", c.Param("id")).Delete(&user)
}

func UpdateUser(c *gin.Context) {
    var user models.User
    config.DB.Where("id = ?", c.Param("id")).First(&user)
    c.BindJSON(&user)
    config.DB.Save(&user)
    c.JSON(200, &user)
}
