package utils

import (
	"strconv"
	"time"
)

func GetYesterdayDate() (int, int, int) {
	yesterday := time.Now().AddDate(0, 0, -1)
	month, err := strconv.Atoi(yesterday.Format("1"))
	HandleError(err)

	year := yesterday.Year()
	day := yesterday.Day()

	return day, month, year
}

func GetYesterdayYmd() string {
	yesterday := time.Now().AddDate(0, 0, -1)
	return yesterday.Format("2006-01-02") // Y-m-d
}
