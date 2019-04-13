package main

import (
	"github.com/Martko/investments-tracker/omaraha"
	"github.com/Martko/investments-tracker/utils"
	"gopkg.in/headzoo/surf.v1"
)

func main() {
	bow := surf.NewBrowser()

	day, month, year := utils.GetYesterdayDate()

	omaraha.FetchAndSaveToDb(bow, day, month, year)
}
