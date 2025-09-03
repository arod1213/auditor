package repo

import (
	"ascap/models"
	"ascap/utils"
	"time"

	"gorm.io/gorm"
)

type CreatePaymentInput struct {
	SongID      uint
	StatementID uint
	Amount      float64
	Quarter     *int
	Year        *int
	Date        *time.Time
}

func CreatePayment(db *gorm.DB, input *CreatePaymentInput) (*models.Payment, error) {
	date := input.Date
	if input.Date == nil {
		date = utils.ParseDate(input.Quarter, input.Year)
	}

	payment := models.Payment{
		Date:        date,
		SongID:      input.SongID,
		StatementID: input.StatementID,
		Amount:      input.Amount,
	}

	err := db.Create(&payment).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

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

func GetStatement(db *gorm.DB, name string) (*models.Statement, error) {
	var statement models.Statement
	err := db.Where("name = ?", name).First(&statement).Error
	if err != nil {
		return nil, err
	}
	return &statement, nil
}
