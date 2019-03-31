package mintos

import (
	"github.com/headzoo/surf/browser"
	"omaraha/utils"
)

func Login(bow *browser.Browser) {
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

func GetPortfolioValues(bow *browser.Browser, currentMonth int, currentYear int) Portfolio {
	// https://www.mintos.com/en/account-statement/?account_statement_filter[fromDate]=01.03.2019&account_statement_filter[toDate]=31.03.2019&account_statement_filter[maxResults]=20

	return Portfolio{
		152.42,
		0,
		152.42,
	}
}

type Portfolio struct {
	InterestAmount float64
	LossAmount     float64
	NetProfit      float64
}
