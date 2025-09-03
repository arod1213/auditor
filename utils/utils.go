package utils

import (
	"time"
)

func ParseDate(quarter *int, year *int) *time.Time {
	if year == nil || quarter == nil {
		return nil
	}

	var month time.Month
	switch *quarter {
	case 1:
		month = time.January
	case 2:
		month = time.April
	case 3:
		month = time.July
	case 4:
		month = time.October
	default:
		month = time.January
	}

	validTime := time.Date(*year, month, 1, 0, 0, 0, 0, time.UTC)
	return &validTime
}
