package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"investments-tracker/utils"
	"log"
)

func GetDbConnection() (db *sql.DB) {
	dbUser, dbPass, dbName := getDatabaseCredentials()

	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@/"+dbName)
	utils.HandleError(err)

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	utils.HandleError(err)

	return db
}

func GetInterestValuesByMonthYear(db *sql.DB, month int, year int) (total, loss, net float64) {
	selectQuery := `
		SELECT
			total,
			loss,
			net
		FROM
			monthly_interests
		WHERE
			source='omaraha' AND 
			month=? AND
			year=?
		LIMIT 1`

	err := db.QueryRow(selectQuery, month, year).Scan(&total, &loss, &net)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, 0, 0
		} else {
			utils.HandleError(err)
		}
	}

	return total, loss, net
}

func InsertValues(values Entry, db *sql.DB) {
	preparedStatement, err := db.Prepare(`
		INSERT INTO 
			daily_interests(date, source, total, loss, net)
			VALUES(?,?,?,?,?)
	`)
	utils.HandleError(err)

	_, _ = preparedStatement.Exec(
		values.Date,
		values.Source,
		values.Total,
		values.Loss,
		values.Net)

	log.Println("inserted to database", values)
}

type Entry struct {
	Date   string
	Source string
	Total  float64
	Loss   float64
	Net    float64
}
