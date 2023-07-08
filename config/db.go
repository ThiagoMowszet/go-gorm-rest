package config

import (
	"github.com/thiagomowszet/go-gorm-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

func Connect() {

    db, err := gorm.Open(postgres.Open("postgres://thiago:Thiago2023@localhost:5432/postgres"), &gorm.Config{})

    if err != nil {
        panic(err)
    }

    db.AutoMigrate(&models.User{})
    DB = db
}
