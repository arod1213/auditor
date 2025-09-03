package models

import (
	"time"
)

type SongtrustPayment struct {
	Description *EmptyString      `csv:"Revenue Class Description"`
	TrackName   string            `csv:"song_name"`
	Amount      float64           `csv:"amount"`
	Date        *yearMonthDayDate `csv:"start_date"`
}

func (a SongtrustPayment) AsPayment() Payment {
	song := Song{
		Name: a.TrackName,
	}
	return Payment{
		RevenueType: Publisher,
		PaymentType: Songtrust,
		Description: a.Description,
		Song:        song,
		Amount:      a.Amount,
		Date:        (*time.Time)(a.Date),
	}
}
