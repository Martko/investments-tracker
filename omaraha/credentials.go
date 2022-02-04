package omaraha

import (
	"os"

	"github.com/Martko/investments-tracker/utils"
	"github.com/joho/godotenv"
)

func getWebPageCredentials() (username, password string) {
	err := godotenv.Load()
	utils.HandleError(err)

	username = os.Getenv("USERNAME")
	password = os.Getenv("PASSWORD")

	return
}
