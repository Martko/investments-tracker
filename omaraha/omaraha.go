package omaraha

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf/browser"
	"investments-tracker/db"
	"investments-tracker/utils"
	"strconv"
	"strings"
)

const (
	currencySeparator = "\u00a0"
	intrestRowNum     = 16 // Intrest incl.
	lossRowNum        = 20 // Principal loss
)

func login(bow *browser.Browser) {
	err := bow.Open("https://omaraha.ee/en/auth/login/")
	utils.HandleError(err)

	fm, err := bow.Form("form.uk-form.uk-form-horizontal")
	utils.HandleError(err)

	username, password := getWebPageCredentials()
	err = fm.Input("username", username)
	utils.HandleError(err)
	err = fm.Input("password", password)
	utils.HandleError(err)
	err = fm.Submit()
	utils.HandleError(err)
}

func getPortfolioValues(bow *browser.Browser, currentMonth int, currentYear int) Portfolio {
	err := bow.Open("https://omaraha.ee/en/invest/stats/")
	utils.HandleError(err)

	var interestAmount, lossAmount, netProfit float64

	rowCount := 0
	bow.Dom().Find("#my_stats tr").Each(func(_ int, s *goquery.Selection) {
		columnCount := 0

		if intrestRowNum == rowCount || lossRowNum == rowCount {
			s.Find("td.uk-text-right.past").Each(func(_ int, e *goquery.Selection) {
				if intrestRowNum == rowCount && columnCount == currentMonth-1 {
					interestAmount = getMonetaryValue(e.Text())
				}
				if lossRowNum == rowCount && columnCount == currentMonth-1 {
					lossAmount = getMonetaryValue(e.Text())
				}
				columnCount++
			})
		}
		rowCount++
	})

	netProfit = interestAmount - lossAmount

	return Portfolio{
		interestAmount,
		lossAmount,
		netProfit,
	}
}

func getMonetaryValue(value string) float64 {
	returnValue, err := strconv.ParseFloat(strings.Split(value, currencySeparator)[0], 64)
	utils.HandleError(err)

	return returnValue
}

func FetchAndSaveToDb(bow *browser.Browser, currentMonth int, currentYear int) {
	login(bow)
	portfolio := getPortfolioValues(bow, currentMonth, currentYear)

	db.InsertOrUpdateDatabase(db.DbEntry{
		"omaraha",
		currentMonth,
		currentYear,
		portfolio.InterestAmount,
		portfolio.LossAmount,
		portfolio.NetProfit,
	})
}

type Portfolio struct {
	InterestAmount float64
	LossAmount     float64
	NetProfit      float64
}
