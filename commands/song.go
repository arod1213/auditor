package commands

import (
	"ascap/repo"
	"fmt"
	"math"
	"strings"

	"gorm.io/gorm"
)

func fetchSong(db *gorm.DB, name string) {
	song, err := repo.GetSong(db, name)
	if err != nil {
		return
	}
	song.ProjectRevenue(db)
	rev := math.Round(song.Revenue(db))
	fmt.Printf("%s -> $%v\n", song.Name, rev)
}

func ReadSong(db *gorm.DB, name string) {
	fmt.Println("name is ", name)
	cleanName := strings.Trim(name, " ")

	if cleanName == "" {
		fmt.Println("please provide a song name")
		return
	}

	fetchSong(db, cleanName)
}
