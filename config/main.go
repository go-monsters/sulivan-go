package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
)

var App = AppConfig{}
var DBConnectionName = DBConnectionNameConfig{}
var DBSqlite = DBSqliteConfig{}

func LoadConfig(){

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env.Parse(&App)
	env.Parse(&DBConnectionName)
	env.Parse(&DBSqlite)
}
