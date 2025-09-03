package models

import (
	"time"
)

type VydiaPayment struct {
	Description *EmptyString   `csv:"Description"`
	TrackName   string         `csv:"Title"`
	Amount      float64        `csv:"USD Amount"`
	Date        *monthYearDate `csv:"Event Date"`
}

func (a VydiaPayment) AsPayment() Payment {
	song := Song{
		Name: a.TrackName,
	}
	return Payment{
		RevenueType: Master,
		PaymentType: Vydia,
		Description: a.Description,
		Song:        song,
		Amount:      a.Amount,
		Date:        (*time.Time)(a.Date),
	}
}
