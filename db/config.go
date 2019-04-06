package db

import (
	"github.com/joho/godotenv"
	"investments-tracker/utils"
	"os"
)

func getDatabaseCredentials() (username, password, database string) {
	err := godotenv.Load()
	utils.HandleError(err)

	username = os.Getenv("DATABASE_USERNAME")
	password = os.Getenv("DATABASE_PASSWORD")
	database = os.Getenv("DATABASE_NAME")

	return
}
