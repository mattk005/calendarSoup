package config

import (
	"time"
)

type DateRange struct {
	StartDate string
	EndDate   string
}

func getDateRange() DateRange {
	// Set today's date
	now := time.Now()

	janFirst := time.Date(
		now.Year(),
		1,
		1,
		0, 1, 0, 0,
		time.Now().Location(),
	)
	janFirstString := janFirst.Format("20060102 150405")

	decThirtyFirst := time.Date(
		now.Year()+1,
		12,
		31,
		23, 59, 0, 0,
		time.Now().Location(),
	)
	decThirtyFirstString := decThirtyFirst.Format("20060102 150405")

	// sixMonths := now.AddDate(0, 6, 0)

	// if now.AddDate(0, 6, 0).After(decThirtyFirst) {
	// 	return janFirstString, sixMonths.Format("20060102 150405")
	// }
	return DateRange{StartDate: janFirstString, EndDate: decThirtyFirstString}
}
