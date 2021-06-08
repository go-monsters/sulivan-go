package main

import (
	"github.com/go-monsters/sulivan/bootstrap"
	"github.com/go-monsters/sulivan/database"
)

func main() {
	bootstrap.Start()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()
	bootstrap.Router.Run()
}
