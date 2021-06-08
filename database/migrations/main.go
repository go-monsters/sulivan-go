package migrations

import (
	"github.com/go-monsters/sulivan/app/models"
	"github.com/go-monsters/sulivan/database"
)

func Migrate() {
	database.DB.AutoMigrate(&models.UserModel{})
}
