package utils

import (
	"strconv"
	"time"
)

func GetCurrentMonth() (int, int) {
	now := time.Now()
	currentMonth, err := strconv.Atoi(now.Format("1"))
	HandleError(err)

	currentYear := now.Year()

	return currentMonth, currentYear
}
