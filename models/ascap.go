package models

import (
	"ascap/utils"
	"fmt"
	"log"
	"time"
)

type AscapPayment struct {
	RevType   RevenueType
	User      *EmptyString      `csv:"Music User"`
	TrackName string            `csv:"Work Title"`
	Dollars   float64           `csv:"Dollars"`
	Amount    float64           `csv:"$ Amount"`
	Quarter   *int              `csv:"Distribution Quarter"`
	Year      *int              `csv:"DistributionYear"`
	Date      *monthDayYearDate `csv:"Distribution Date"`
}

func (a AscapPayment) AsPayment() Payment {
	song := Song{
		Name: a.TrackName,
	}

	var date *time.Time
	if a.Date != nil {
		date = (*time.Time)(a.Date)
	} else {
		date = utils.ParseDate(a.Quarter, a.Year)
		if date == nil {
			log.Printf("date is nil - %v %v\n", a.Quarter, a.Year)
		} else {
			fmt.Printf("new date is %v\n", *date)
		}
	}

	return Payment{
		RevenueType: a.RevType,
		PaymentType: ASCAP,
		Description: a.User,
		Song:        song,
		Amount:      a.Dollars + a.Amount,
		Date:        date,
	}
}
