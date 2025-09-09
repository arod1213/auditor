package models

import (
	"fmt"
	"math"
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	SongID      uint         `gorm:"not null" json:"song_id"`
	Song        Song         `gorm:"foreignKey:SongID" json:"song"`
	Amount      float64      `gorm:"not null" csv:"Dollars"`
	Description *EmptyString `csv:"Description"`
	StatementID uint         `gorm:"not null" json:"statement_id"`
	Statement   Statement    `gorm:"foreignKey:StatementID" json:"statement"`
	Date        *time.Time   `json:"date"`
	PaymentType PaymentType  `json:"payment_type"`
	RevenueType RevenueType  `json:"revenue_type"`
}

func (p Payment) String() string {
	dateStr := "nil"
	if p.Date != nil {
		dateStr = p.Date.Format("2006-01-02")
	}

	amount := math.Round(p.Amount*100) / 100
	return fmt.Sprintf("Title: %s -> $: %v at %s\n", p.Song.Name, amount, dateStr)
}

type Royalty interface {
	AsPayment() *Payment
}

type PaymentType int

const (
	ASCAP PaymentType = iota
	Vydia
	MLC
	SoundExchange
	Distrokid
	Songtrust
)

var strToPaymentType = map[string]PaymentType{
	"ascap":          ASCAP,
	"vydia":          Vydia,
	"mlc":            MLC,
	"sound exchange": SoundExchange,
	"distrokid":      Distrokid,
	"songtrust":      Songtrust,
}

func PaymentTypeFromString(s string) (PaymentType, error) {
	if rt, ok := strToPaymentType[s]; ok {
		return rt, nil
	}
	return 0, fmt.Errorf("invalid PaymentType: %s", s)
}

type RevenueType int

const (
	Publisher RevenueType = iota
	Writer
	Master
)
