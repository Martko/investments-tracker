package main

import (
	"gopkg.in/headzoo/surf.v1"
	"investments-tracker/omaraha"
	"investments-tracker/utils"
)

func main() {
	bow := surf.NewBrowser()

	currentDay, currentMonth, currentYear := utils.GetCurrentDate()

	omaraha.FetchAndSaveToDb(bow, currentDay, currentMonth, currentYear)
}
