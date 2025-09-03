package models

import (
	"time"
)

type MLCPayment struct {
	UseType   *EmptyString      `csv:"Use Type"`
	TrackName string            `csv:"Work Primary Title"`
	Amount    float64           `csv:"Distributed Amount"`
	Date      *yearMonthDayDate `csv:"Distribution Date"`
}

func (a MLCPayment) AsPayment() Payment {
	song := Song{
		Name: a.TrackName,
	}

	return Payment{
		RevenueType: Publisher,
		PaymentType: MLC,
		Description: a.UseType,
		Song:        song,
		Amount:      a.Amount,
		Date:        (*time.Time)(a.Date),
	}
}
