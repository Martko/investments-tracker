package omaraha

import (
	"github.com/joho/godotenv"
	"investments-tracker/utils"
	"os"
)

func getWebPageCredentials() (username, password string) {
	err := godotenv.Load()
	utils.HandleError(err)

	username = os.Getenv("USERNAME")
	password = os.Getenv("PASSWORD")

	return
}
