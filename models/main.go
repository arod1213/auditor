package models

import (
	"strings"

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

type Statement struct {
	gorm.Model
	Name string `gorm:"not null;unique" json:"name"`
}
