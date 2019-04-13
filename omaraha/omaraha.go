package omaraha

import (
	"github.com/Martko/investments-tracker/db"
	"github.com/Martko/investments-tracker/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf/browser"
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

func getInterestValues(bow *browser.Browser, currentMonth int) Portfolio {
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

func FetchAndSaveToDb(bow *browser.Browser, currentDay int, currentMonth int, currentYear int) {
	login(bow)
	portfolio := getInterestValues(bow, currentMonth)
	connection := db.GetDbConnection()

	mTotal, mLoss, mNet := db.GetInterestValuesByMonthYear(connection, currentMonth, currentYear)

	total := portfolio.Total - mTotal
	loss := portfolio.Loss - mLoss
	net := portfolio.Net - mNet

	db.InsertValues(db.Entry{
		Date:   utils.GetYesterdayYmd(),
		Source: "omaraha",
		Total:  total,
		Loss:   loss,
		Net:    net,
	}, connection)
}

type Portfolio struct {
	Total float64
	Loss  float64
	Net   float64
}
