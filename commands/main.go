package commands

import (
	"ascap/models"
	"ascap/reader"
	"ascap/repo"
	"fmt"
	"os"

	"gorm.io/gorm"
)

func readAll(db *gorm.DB) {
	path := os.Getenv("VYDIA_PATH")
	reader.ReadStatements(db, path, reader.ReadVydia)
	path = os.Getenv("DISTROKID_PATH")
	reader.ReadStatements(db, path, reader.ReadDistrokid)
	path = os.Getenv("MLC_PATH")
	reader.ReadStatements(db, path, reader.ReadMLC)
	path = os.Getenv("ASCAP_PATH")
	reader.ReadStatements(db, path, reader.ReadASCAP)
	path = os.Getenv("SX_PATH")
	reader.ReadStatements(db, path, reader.ReadSX)
	path = os.Getenv("SONGTRUST_PATH")
	reader.ReadStatements(db, path, reader.ReadSongtrust)
}

func fetchSongData(db *gorm.DB, name string, pt *models.PaymentType) {
	song, err := repo.GetSong(db, name)
	if err != nil {
		return
	}
	GetSongData(db, song, pt)
}

func monthly(db *gorm.DB) {
	sum, _ := GetMonthlyTotals(db)
	for _, s := range sum {
		fmt.Print(s)
	}
}
