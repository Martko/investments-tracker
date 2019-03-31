package mintos

import (
	"github.com/joho/godotenv"
	"investments-tracker/utils"
	"os"
)

func getWebPageCredentials() (username, password string) {
	err := godotenv.Load()
	utils.HandleError(err)

	username = os.Getenv("MINTOS_USERNAME")
	password = os.Getenv("MINTOS_PASSWORD")

	return
}
