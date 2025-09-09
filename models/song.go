package models

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"gorm.io/gorm"
)

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
