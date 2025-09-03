package models

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"gorm.io/gorm"
)

type EmptyString string

func (a *EmptyString) UnmarshalCSV(s string) error {
	cleaned := strings.TrimSpace(s)
	if cleaned == "" {
		return nil
	}
	*a = EmptyString(cleaned)
	return nil
}

type Royalty interface {
	AsPayment() *Payment
}

type Statement struct {
	gorm.Model
	Name string `gorm:"not null;unique" json:"name"`
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

type Song struct {
	gorm.Model
	Name string `gorm:"not null;unique"`
}

func (s *Song) BeforeSave(tx *gorm.DB) error {
	s.Name = strings.ToLower(strings.TrimSpace(s.Name))
	return nil
}

func (s *Song) Revenue(db *gorm.DB) float64 {
	var total float64
	err := db.
		Model(&Payment{}).
		Where("song_id = ? AND amount > 0", s.ID).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&total).Error

	if err != nil {
		return 0
	}
	return total
}

type monthlyRev struct {
	Date  time.Time
	Total float64
}

func (s Song) ProjectRevenue(db *gorm.DB) float64 {
	monthlyRev, err := s.MonthlyRevenue(db)
	if err != nil {
		return 0
	}

	data := *monthlyRev
	sort.Slice(data, func(i, j int) bool {
		return data[i].Date.After(data[j].Date)
	})

	for _, val := range data {
		fmt.Printf("%v -> $%v\n", val.Date, val.Total)
	}

	return 0
}

func (s Song) MonthlyRevenue(db *gorm.DB) (*[]monthlyRev, error) {
	var res []monthlyRev

	err := db.
		Model(Payment{}).
		Select("date, strftime('%Y-%m-%d', date) as month, COALESCE(SUM(amount), 0) as total").
		Where("song_id = ?", s.ID).
		Group("month").
		Order("month").
		Scan(&res).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &res, nil
}
