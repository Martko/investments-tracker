package omaraha

import (
	"github.com/Martko/investments-tracker/utils"
	"github.com/joho/godotenv"
	"os"
)

func getWebPageCredentials() (username, password string) {
	err := godotenv.Load()
	utils.HandleError(err)

	username = os.Getenv("USERNAME")
	password = os.Getenv("PASSWORD")

	return
}
