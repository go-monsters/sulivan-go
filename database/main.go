package database

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	//db.LogMode(true)
	DB = db
	return DB
}

