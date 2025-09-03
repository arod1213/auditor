package commands

import (
	"ascap/models"
	"fmt"
	"math"
	"slices"
	"strings"

	"gorm.io/gorm"
)

func fetchPlatform(db *gorm.DB, paymentType models.PaymentType) {
	var reg = []string{
		"about time", "am i even living anymore", "baby tee", "bird's eye view",
		"bird s eye view", "body bag", "bodybag", "sunroof (feat. manuel turizo)", "sunroof (loud luxury remix)", "sunroof 24 kgoldn remix", "call my bluff", "crying over", "february 7th.",
		"felt too good!", "fool for you (amor en vano)", "gentle hellraiser", "ground",
		"growing up is _____", "hands on my body", "honeymoon", "hopeless romantic",
		"kiss me like a secret", "magic", "never catch a break", "one of a kind",
		"on the day i leave", "persimmon", "queen's disease", "sexy villain",
		"trip around the sun", "used to be yours", "vienna (in memoriam)",
		"where you been hiding", "wild again", "wish i had you", "drunk tank", "crying over you",
		"strawberry blonde", "sunroof (manuel turizo remix)", "sunroof (thomas rhett remix)",
		"halfway", "sunroof", "sunroof (24kgoldn remix)", "used to be", "sexy villian",
		"fuckboys", "fckboys", "cynical (put on my boots)", "fools love", "let down your walls",
		"never catch a break (feat. marc e. bassy)", "timeless", "fool for you (amor en vano)",
		"lifestyle (feat. dan raff)", "3am", "home", "demure", "favorite daughter", "over reliance",
		"solo", "perfect slumber", "snapshot (feat. ctl)", "stop saying sorry",
	}

	songs := RegSongs(db, paymentType)
	var sum float64
	for songName, payout := range songs {
		if slices.Contains(reg, songName) {
			sum += payout
			fmt.Printf("%v -> $%v\n", songName, math.Round(payout))
			continue
		}
		fmt.Printf("Missing -> %v\n", songName)
	}
	fmt.Printf("Total -> $%v\n", math.Round(sum))
}

func ReadPlatform(db *gorm.DB, name string) {
	cleanName := strings.Trim(name, " ")

	if cleanName == "" {
		fmt.Println("please provide a platform name")
		return
	}

	paymentType, err := models.PaymentTypeFromString(strings.ToLower(cleanName))
	if err != nil {
		fmt.Println(err)
		return
	}
	fetchPlatform(db, paymentType)
}
