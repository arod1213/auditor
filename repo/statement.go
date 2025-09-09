package repo

import (
	"ascap/models"

	"gorm.io/gorm"
)

func CreateStatement(db *gorm.DB, name string) (*models.Statement, error) {
	statement := models.Statement{
		Name: name,
	}
	err := db.Create(&statement).Error
	if err != nil {
		return nil, err
	}
	return &statement, nil
}

func GetStatement(db *gorm.DB, name string) (*models.Statement, error) {
	var statement models.Statement
	err := db.Where("name = ?", name).First(&statement).Error
	if err != nil {
		return nil, err
	}
	return &statement, nil
}
