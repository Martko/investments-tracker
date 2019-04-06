package main

import (
	"gopkg.in/headzoo/surf.v1"
	"investments-tracker/funderbeam"
	"investments-tracker/mintos"
	"investments-tracker/omaraha"
	"investments-tracker/utils"
)

func main() {
	bow := surf.NewBrowser()

	currentMonth, currentYear := utils.GetCurrentMonth()

	mintos.FetchAndSaveToDb(bow, currentMonth, currentYear)
	omaraha.FetchAndSaveToDb(bow, currentMonth, currentYear)
}
