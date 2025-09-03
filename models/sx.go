package models

import (
	"time"
)

type SXPayment struct {
	Description *EmptyString      `csv:"Category of Service"`
	TrackName   string            `csv:"Track Name"`
	Amount      float64           `csv:"Your Payment Amount"`
	Date        *dayMonthYearDate `csv:"Broadcast End Date"`
}

func (a SXPayment) AsPayment() Payment {
	song := Song{
		Name: a.TrackName,
	}
	return Payment{
		RevenueType: Master,
		PaymentType: SoundExchange,
		Description: a.Description,
		Song:        song,
		Amount:      a.Amount,
		Date:        (*time.Time)(a.Date),
	}
}
