package reader

import (
	"ascap/models"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

func ReadSongtrust(path string) (*[]models.Payment, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()

	var payments []models.SongtrustPayment
	if err := gocsv.Unmarshal(file, &payments); err != nil {
		fmt.Println("Error unmarshaling CSV:", err)
		return nil, err
	}

	var royalties []models.Payment
	for _, p := range payments {
		royalties = append(royalties, p.AsPayment())
	}
	return &royalties, nil
}
