package main

import (
	"gopkg.in/headzoo/surf.v1"
	"investments-tracker/omaraha"
	"investments-tracker/utils"
)

func main() {
	bow := surf.NewBrowser()

	day, month, year := utils.GetYesterdayDate()

	omaraha.FetchAndSaveToDb(bow, day, month, year)
}
