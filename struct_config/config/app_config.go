package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBParams struct {
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Address  string `mapstructure:"address"`
		Database string `mapstructure:"database"`
		Charset  string `mapstructure:"charset"`
	} `mapstructure:"db_params"`
	ServerAddr string `mapstructure:"server_address"`
}

var AppConfig Config

func LoadConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	viper.Unmarshal(&AppConfig)
}
