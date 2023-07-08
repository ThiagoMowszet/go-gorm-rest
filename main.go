package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thiagomowszet/go-gorm-rest/config"
	r "github.com/thiagomowszet/go-gorm-rest/routes"
)

func main() {
    router := gin.New()
    config.Connect()

    r.UserRoute(router)

    // Server listen in port :3030
    router.Run(":3000")
}

