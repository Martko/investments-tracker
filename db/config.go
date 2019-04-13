package db

import (
	"github.com/Martko/investments-tracker/utils"
	"github.com/joho/godotenv"
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
