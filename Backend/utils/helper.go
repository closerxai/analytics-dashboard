package utils

import "time"

func ApplyDefaultMonth(startDate, endDate string) (string, string) {
	// If user provided both → use them
	if startDate != "" && endDate != "" {
		return startDate, endDate
	}

	now := time.Now().UTC()
	first := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	last := first.AddDate(0, 1, 0).Add(-time.Second)

	return first.Format("2006-01-02"), last.Format("2006-01-02")
}


func Abs64(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}