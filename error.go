package date

import (
	"errors"
	"time"
)

var (
	ErrLessFirstDay = errors.New("date: less first Julian day")
	ErrZeroYear     = errors.New("date: no year 0 in the Julian calendar")
	ErrPassage      = errors.New("date: passage from Julian to Gregorian calendar")
	ErrInvalidMonth = errors.New("date: invalid month")
	ErrInvalidDay   = errors.New("date: invalid day")
)

func checkDate(year int, month time.Month, day int) error {

	// Check FirstJulianDay
	if year < firstJulianDay.year {
		return ErrLessFirstDay
	}
	if year == firstJulianDay.year {
		if month < firstJulianDay.month {
			return ErrLessFirstDay
		}
		if month == firstJulianDay.month {
			if day < firstJulianDay.day {
				return ErrLessFirstDay
			}
		}
	}

	// there is no year 0 in the Julian calendar
	if year == 0 {
		return ErrZeroYear
	}

	// passage from Julian to Gregorian calendar
	if (year == passageYear) && (month == passageMonth) &&
		((day > lastJulianDay.day) && (day < firstGregorianDay.day)) {
		return ErrPassage
	}

	// Check month
	if (month < time.January) || (month > time.December) {
		return ErrInvalidMonth
	}

	// Check day
	if (day < 1) || (day > NumberOfDays(year, month)) {
		return ErrInvalidDay
	}

	return nil
}
