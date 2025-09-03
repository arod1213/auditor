package models

import (
	"fmt"
	"strings"
	"time"
)

type monthDayYearDate time.Time

func (cd *monthDayYearDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse("01-02-2006", s)
	if err != nil {
		return err
	}
	*cd = monthDayYearDate(t)
	return nil
}

func (cd *monthDayYearDate) UnmarshalCSV(s string) error {
	if s == "" {
		return nil
	}
	t, err := time.Parse("01-02-2006", s) // Adjust format if needed
	if err != nil {
		return fmt.Errorf("failed to parse date '%s': %w", s, err)
	}
	*cd = monthDayYearDate(t)
	return nil
}

type yearMonthDayDate time.Time

func (cd *yearMonthDayDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*cd = yearMonthDayDate(t)
	return nil
}

func (cd *yearMonthDayDate) UnmarshalCSV(s string) error {
	if s == "" {
		return nil
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return fmt.Errorf("failed to parse date '%s': %w", s, err)
	}
	*cd = yearMonthDayDate(t)
	return nil
}

type dayMonthYearDate time.Time

func (cd *dayMonthYearDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse("02-Jan-2006", s)
	if err != nil {
		return err
	}
	*cd = dayMonthYearDate(t)
	return nil
}

func (cd *dayMonthYearDate) UnmarshalCSV(s string) error {
	if s == "" {
		return nil
	}
	t, err := time.Parse("02-Jan-2006", s)
	if err != nil {
		return fmt.Errorf("failed to parse date '%s': %w", s, err)
	}
	*cd = dayMonthYearDate(t)
	return nil
}

type monthYearDate time.Time

func (cd *monthYearDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse("01/2006", s)
	if err != nil {
		return err
	}
	*cd = monthYearDate(t)
	return nil
}

func (cd *monthYearDate) UnmarshalCSV(s string) error {
	if s == "" {
		return nil
	}
	t, err := time.Parse("01/2006", s)
	if err != nil {
		return fmt.Errorf("failed to parse date '%s': %w", s, err)
	}
	*cd = monthYearDate(t)
	return nil
}
