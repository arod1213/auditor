package repo

import (
	"ascap/models"

	"gorm.io/gorm"
)

func CreateSong(db *gorm.DB, name string) (*models.Song, error) {
	song := models.Song{
		Name: name,
	}
	err := db.Create(&song).Error
	if err != nil {
		return nil, err
	}
	return &song, nil
}

func CreateOrFindSong(db *gorm.DB, name string) (*models.Song, error) {
	song, err := GetSong(db, name)
	if err == nil {
		return song, nil
	}

	song, err = CreateSong(db, name)
	if err != nil {
		return nil, err
	}
	return song, err
}

func GetSong(db *gorm.DB, name string) (*models.Song, error) {
	var song models.Song
	err := db.
		Where("LOWER(TRIM(name)) = LOWER(TRIM(?))", name).
		First(&song).Error
	if err != nil {
		return nil, err
	}
	return &song, nil
}
