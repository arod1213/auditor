package reader

import (
	"ascap/models"
	"ascap/repo"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

func GetStatementNames(path string) ([]string, error) {
	var csvs []string = make([]string, 0)

	dir, err := os.ReadDir(path)
	if err != nil {
		return csvs, err
	}

	for _, file := range dir {
		ext := filepath.Ext(file.Name())
		if ext != ".csv" {
			continue
		}
		// fullpath := filepath.Join(path, file.Name())
		csvs = append(csvs, file.Name())
	}
	return csvs, nil
}

func ReadStatements(db *gorm.DB, path string, reader func(path string) (*[]models.Payment, error)) error {
	statements, err := GetStatementNames(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var totals map[string]float64 = make(map[string]float64)
	for _, path := range statements {

		newStatement, err := repo.CreateStatement(db, path)
		if err != nil {
			fmt.Print(err)
			continue
		}

		data, err := reader(path)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, entry := range *data {
			err := SavePayment(db, &entry, newStatement)
			if err != nil {
				fmt.Println(err)
				continue
			}
			totals[entry.Song.Name] += entry.Amount
		}
		// fmt.Print(data)
	}
	fmt.Print(totals)
	return nil
}

func SavePayment(db *gorm.DB, payment *models.Payment, statement *models.Statement) error {
	song, err := repo.CreateOrFindSong(db, payment.Song.Name)
	if err != nil {
		return err
	}
	payment.Song = *song
	payment.Statement = *statement
	err = db.Save(&payment).Error
	return nil
}

func Print(data any) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ") // Optional: makes it pretty
	encoder.Encode(data)
}
