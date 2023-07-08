package routes

import (
	"github.com/gin-gonic/gin"
	c "github.com/thiagomowszet/go-gorm-rest/controller"
)


func UserRoute(router *gin.Engine) {
    router.GET("/", c.GetUsers)
    router.POST("/", c.CreateUser)
    router.DELETE("/:id", c.DeleteUser)
    router.PUT("/:id", c.UpdateUser)
}
