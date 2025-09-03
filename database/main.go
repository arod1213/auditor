package database

import (
	"ascap/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func EstablishConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./info.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&models.Payment{}, &models.Song{}, &models.Statement{})
	return db, nil
}
