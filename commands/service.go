package commands

import (
	"ascap/models"
	"fmt"
	"log"
	"math"

	"gorm.io/gorm"
)

func GetSum(db *gorm.DB, name string) float64 {
	var song models.Song
	err := db.Where("name = ?", name).First(&song).Error
	if err != nil {
		return 0
	}
	var payments []models.Payment
	err = db.Where("song_id = ?", song.ID).Find(&payments).Error
	if err != nil {
		return 0
	}

	var sum float64
	for _, s := range payments {
		sum += s.Amount
	}
	return sum
}

func GetStatementTotals(db *gorm.DB) []float64 {
	var totals []float64 = make([]float64, 0)

	var statements []models.Statement
	err := db.Find(&statements).Error
	if err != nil {
		log.Print(err)
		return totals
	}

	for _, s := range statements {
		var total float64
		err := db.Model(&models.Payment{}).
			Select("COALESCE(SUM(amount), 0)").
			Having("SUM(amount) >= 0").
			Where("statement_id = ?", s.ID).
			Scan(&total).Error

		if err != nil {
			log.Print(err)
			continue
		}

		log.Printf("Statement %s -> Total %v\n", s.Name, total)

		totals = append(totals, total)
	}
	return totals
}

type MonthlyTotal struct {
	Month string // e.g., "2025-08"
	Total float64
}

func (m MonthlyTotal) String() string {
	return fmt.Sprintf("Month %s -> $%v\n", m.Month, m.Total)
}

func GetMonthlyTotals(db *gorm.DB) ([]MonthlyTotal, error) {
	var results []MonthlyTotal

	err := db.
		Model(&models.Payment{}).
		Preload("Song").
		Select("strftime('%Y-%m', date) as month, COALESCE(SUM(amount), 0) as total").
		Group("month").
		Having("SUM(amount) >= 0").
		Order("month").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}

func GetSongTotals(db *gorm.DB) {
	var songs []models.Song
	err := db.Find(&songs).Error
	if err != nil {
		return
	}
	for _, s := range songs {
		rev := s.Revenue(db)
		fmt.Printf("%s -> %v\n", s.Name, rev)
	}
}

func RegSongs(db *gorm.DB, pt models.PaymentType) map[string]float64 {
	var songNames map[string]float64 = make(map[string]float64, 100)
	var payments []models.Payment
	err := db.Preload("Song").Where("payment_type = ?", pt).Find(&payments).Error
	if err != nil {
		fmt.Print(err)
		return songNames
	}

	for _, s := range payments {
		songNames[s.Song.Name] += s.Amount
	}

	return songNames
}

func GetSongData(db *gorm.DB, song *models.Song, pt *models.PaymentType) {
	var payments []models.Payment

	query := db.
		Select("strftime('%Y-%m', date) as month, payment_type, song_id, COALESCE(SUM(amount), 0) as amount").
		Preload("Song").
		Where("payments.song_id = ?", song.ID).
		Group("strftime('%Y-%m', date), payment_type, song_id")

	if pt != nil {
		query = query.Where("payments.payment_type = ?", pt)
	}

	err := query.Find(&payments).Error

	if err != nil {
		return
	}
	var sum float64
	for _, p := range payments {
		fmt.Print(p)
		sum += p.Amount
	}
	fmt.Printf("Total -> $%v\n", math.Round(sum*100)/100)

}
