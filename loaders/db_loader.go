package loaders

import (
	"Mohamed-Abdelrazeq/o-auth-2/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseInstance struct {
	DB *gorm.DB
}

func LoadDB() *DatabaseInstance {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	// Init DB Instance
	return &DatabaseInstance{db}
}
