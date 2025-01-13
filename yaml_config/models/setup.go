package models

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db_params := viper.GetStringMapString("db_params")
	database, err := gorm.Open(mysql.Open(db_params["username"]+":"+db_params["password"]+"@tcp("+db_params["address"]+")/"+db_params["database"]+"?charset="+db_params["charset"]+"&parseTime=true"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}

func DBMigrate() {
	DB.AutoMigrate(&Blog{})
}
