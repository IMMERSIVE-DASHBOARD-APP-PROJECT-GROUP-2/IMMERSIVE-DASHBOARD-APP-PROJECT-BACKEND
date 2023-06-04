package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	SECRET_JWT = ""
)

type AppConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOSTNAME string
	DB_PORT     int
	DB_NAME     string
	jwtKey      string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {

	app := AppConfig{}

	viper.AddConfigPath(".")
	viper.SetConfigName("local")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("error read config:", err.Error())
		return nil
	}

	app.jwtKey = viper.GetString("JWT_KEY")
	app.DB_USERNAME = viper.GetString("DBUSER")
	app.DB_PASSWORD = viper.GetString("DBPASS")
	app.DB_HOSTNAME = viper.GetString("DBHOST")
	app.DB_PORT = viper.GetInt("DBPORT")
	app.DB_NAME = viper.GetString("DBNAME")

	SECRET_JWT = app.jwtKey

	return &app
}
