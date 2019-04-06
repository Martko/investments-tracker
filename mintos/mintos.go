package mintos

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf/browser"
	"investments-tracker/db"
	"investments-tracker/utils"
	"strconv"
)

const (
	intrestRowNum        = 5 // Interest received
	intrestOnRebuyRowNum = 7 // Interest income on rebuy
)

func login(bow *browser.Browser) {
	err := bow.Open("https://www.mintos.com/en/login")
	utils.HandleError(err)

	fm, err := bow.Form("#login-form")
	utils.HandleError(err)

	username, password := getWebPageCredentials()
	err = fm.Input("_username", username)
	utils.HandleError(err)
	err = fm.Input("_password", password)
	utils.HandleError(err)
	err = fm.Submit()
	utils.HandleError(err)
}

func getPortfolioValues(bow *browser.Browser, currentMonth int, currentYear int) Portfolio {
	/*
		currentMonthString := strconv.Itoa(currentMonth)
		currentYearString := strconv.Itoa(currentYear)

		fromDate := "01." + currentMonthString + "." + currentYearString
		toDate := "31." + currentMonthString + "." + currentYearString

	*/
	fromDate := "01.03.2019"
	toDate := "31.03.2019"

	err := bow.Open("https://www.mintos.com/en/account-statement/?" +
		"account_statement_filter[fromDate]=" + fromDate + "" +
		"&account_statement_filter[toDate]=" + toDate + "" +
		"&account_statement_filter[maxResults]=20")

	utils.HandleError(err)

	var interestAmount float64

	rowCount := 0
	bow.Dom().Find("#overview-results table tbody tr").Each(func(_ int, s *goquery.Selection) {
		if intrestRowNum == rowCount || intrestOnRebuyRowNum == rowCount {
			interestAmount += getMonetaryValue(s.Find("span.mod-pointer").Text())
		}
		rowCount++
	})

	return Portfolio{interestAmount, 0, interestAmount}
}

func FetchAndSaveToDb(bow *browser.Browser, currentMonth int, currentYear int) {
	login(bow)
	portfolio := getPortfolioValues(bow, currentMonth, currentYear)

	db.InsertOrUpdateDatabase(db.DbEntry{
		"mintos",
		currentMonth,
		currentYear,
		portfolio.InterestAmount,
		portfolio.LossAmount,
		portfolio.NetProfit,
	})
}

func getMonetaryValue(value string) float64 {
	returnValue, err := strconv.ParseFloat(value, 64)
	utils.HandleError(err)

	return returnValue
}

type Portfolio struct {
	InterestAmount float64
	LossAmount     float64
	NetProfit      float64
}
