package main

import (
	"github.com/headzoo/surf/browser"
	"gopkg.in/headzoo/surf.v1"
	"investments-tracker/mintos"
	"investments-tracker/omaraha"
)

func main() {
	bow := surf.NewBrowser()

	currentMonth, currentYear := getCurrentMonth()

	//handleMintos(bow, currentMonth, currentYear)
	handleOmaraha(bow, currentMonth, currentYear)

}

func handleOmaraha(bow *browser.Browser, currentMonth int, currentYear int) {
	omaraha.Login(bow)
	portfolio := omaraha.GetPortfolioValues(bow, currentMonth, currentYear)

	insertOrUpdateDatabase(dbEntry{
		"omaraha",
		currentMonth,
		currentYear,
		portfolio.InterestAmount,
		portfolio.LossAmount,
		portfolio.NetProfit,
	})
}

func handleMintos(bow *browser.Browser, currentMonth int, currentYear int) {
	mintos.Login(bow)
	portfolio := mintos.GetPortfolioValues(bow, currentMonth, currentYear)

	insertOrUpdateDatabase(dbEntry{
		"mintos",
		currentMonth,
		currentYear,
		portfolio.InterestAmount,
		portfolio.LossAmount,
		portfolio.NetProfit,
	})
}
