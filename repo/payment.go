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
