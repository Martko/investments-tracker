package utils

import (
	"strconv"
	"time"
)

func GetCurrentDate() (int, int, int) {
	now := time.Now()
	currentMonth, err := strconv.Atoi(now.Format("1"))
	HandleError(err)

	currentYear := now.Year()
	currentDay := now.Day()

	return currentDay, currentMonth, currentYear
}
