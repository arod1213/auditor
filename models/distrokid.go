package models

import (
	"time"
)

type DistrokidPayment struct {
	Store     *EmptyString      `csv:"Store"`
	TrackName string            `csv:"Title"`
	Amount    float64           `csv:"Earnings (USD)"`
	Date      *yearMonthDayDate `csv:"Reporting Date"`
}

func (a DistrokidPayment) AsPayment() Payment {
	song := Song{
		Name: a.TrackName,
	}

	return Payment{
		RevenueType: Master,
		PaymentType: Distrokid,
		Description: a.Store,
		Song:        song,
		Amount:      a.Amount,
		Date:        (*time.Time)(a.Date),
	}
}
