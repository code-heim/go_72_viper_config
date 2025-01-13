package models

import (
	"go_blogs/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open(config.AppConfig.DBParams.Username+
		":"+
		config.AppConfig.DBParams.Password+
		"@tcp("+config.AppConfig.DBParams.Address+")/"+
		config.AppConfig.DBParams.Database+
		"?charset="+
		config.AppConfig.DBParams.Charset+
		"&parseTime=true"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}

func DBMigrate() {
	DB.AutoMigrate(&Blog{})
}
