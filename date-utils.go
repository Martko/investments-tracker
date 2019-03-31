package main

import (
	"omaraha/utils"
	"strconv"
	"time"
)

func getCurrentMonth() (int, int) {
	now := time.Now()
	currentMonth, err := strconv.Atoi(now.Format("1"))
	utils.HandleError(err)

	currentYear := now.Year()

	return currentMonth, currentYear
}
