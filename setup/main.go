package setup

import (
	"ascap/models"
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		fmt.Print("env could not be loaded")
		return err
	}
	return nil
}

func EstablishConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./info.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&models.Payment{}, &models.Song{}, &models.Statement{})
	return db, nil
}
