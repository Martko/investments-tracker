package omaraha

import (
	"strconv"
	"strings"

	"github.com/Martko/investments-tracker/db"
	"github.com/Martko/investments-tracker/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf/browser"
)

const (
	currencySeparator       = "\u00a0"
	intrestRowNum           = 17 // Intrest incl.
	lossRowNum              = 21 // Principal loss
	portfolioValueRowNum    = 5  // Portfolio balance
	initialInvestmentRowNum = 1  // Money in the portal
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

func getInterestValuesFromStatistics(bow *browser.Browser, currentMonth int) Portfolio {
	err := bow.Open("https://omaraha.ee/en/invest/stats/")
	utils.HandleError(err)

	var interestAmount, lossAmount, netProfit, portfolioValue, initialInvestment float64

	rowCount := 0
	bow.Dom().Find("#my_stats tr").Each(func(_ int, s *goquery.Selection) {
		columnCount := 0

		if intrestRowNum == rowCount || lossRowNum == rowCount || portfolioValueRowNum == rowCount {
			s.Find("td.uk-text-right.past").Each(func(_ int, e *goquery.Selection) {
				if intrestRowNum == rowCount && columnCount == currentMonth-1 {
					interestAmount = getMonetaryValue(e.Text())
				}
				if lossRowNum == rowCount && columnCount == currentMonth-1 {
					lossAmount = getMonetaryValue(e.Text())
				}
				if portfolioValueRowNum == rowCount && columnCount == currentMonth-1 {
					portfolioValue = getMonetaryValue(e.Text())
				}
				columnCount++
			})
		}

		if initialInvestmentRowNum == rowCount {
			s.Find("td.uk-text-right.cumulative").Each(func(_ int, e *goquery.Selection) {
				initialInvestment = getMonetaryValue(e.Text())
			})
		}
		rowCount++
	})

	netProfit = interestAmount - lossAmount

	return Portfolio{
		interestAmount,
		lossAmount,
		netProfit,
		portfolioValue,
		initialInvestment,
	}
}

func getMonetaryValue(value string) float64 {
	value = strings.Replace(value, " ", "", -1)
	value = strings.Replace(value, "\n", "", -1)
	value = strings.Replace(value, ",", "", -1)
	returnValue, err := strconv.ParseFloat(strings.Split(value, currencySeparator)[0], 64)
	utils.HandleError(err)

	return returnValue
}

func getAvailableMoney(bow *browser.Browser) float64 {
	rowCount := 0
	var availableMoney float64

	bow.Dom().Find(".uk-width-medium-2-3 table tbody tr").Each(func(_ int, s *goquery.Selection) {
		columnCount := 0

		s.Find("td").Each(func(_ int, e *goquery.Selection) {
			if columnCount == 1 {
				value := strings.Replace(e.Text(), ",", ".", -1)
				availableMoney += getMonetaryValue(value)
			}
			columnCount++
		})
		rowCount++
	})

	return availableMoney
}

// FetchAndSaveToDb fetches data from the page and store in DB
func FetchAndSaveToDb(bow *browser.Browser, currentDay int, currentMonth int, currentYear int) {
	login(bow)

	availableMoney := getAvailableMoney(bow)
	portfolio := getInterestValuesFromStatistics(bow, currentMonth)

	connection := db.GetDbConnection()

	mTotal, mLoss, mNet := db.GetInterestValuesByMonthYear(connection, currentMonth, currentYear)

	total := portfolio.Total - mTotal
	loss := portfolio.Loss - mLoss
	net := portfolio.Net - mNet

	db.InsertInterestValues(db.Entry{
		Date:       utils.GetYesterdayYmd(),
		Source:     "omaraha",
		AssetClass: "loans_partially_secured",
		Total:      total,
		Loss:       loss,
		Net:        net,
	}, connection)

	db.InsertPortfolioValues(db.PortfolioValueEntry{
		Date:              utils.GetYesterdayYmd(),
		Source:            "omaraha",
		Value:             portfolio.PortfolioValue,
		InitialInvestment: portfolio.InitialInvestment,
		Profit:            portfolio.PortfolioValue - portfolio.InitialInvestment,
		Cash:              availableMoney,
	}, connection)
}

type Portfolio struct {
	Total             float64
	Loss              float64
	Net               float64
	PortfolioValue    float64
	InitialInvestment float64
}
