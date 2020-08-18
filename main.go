package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"mp_royal/configs"
	"mp_royal/models"
)

func main() {
	var err error
	r := gin.Default()

	configs.DB, err = gorm.Open(mysql.Open(configs.DbURL(configs.BuildDBConfig())), &gorm.Config{
		//Logger: newLogger,
	})

	if err != nil {
		log.Println(err)
	}

	configs.DB.AutoMigrate(&models.HairRelease{}, models.FaceRelease{}, &models.MonthlyFace{}, &models.MonthlyHair{})

	r.Run(":8080")
}


