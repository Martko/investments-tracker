package db

import (
	"database/sql"
	"log"

	"github.com/Martko/investments-tracker/utils"
	_ "github.com/go-sql-driver/mysql"
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
			monthly_passive_income
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

func InsertInterestValues(values Entry, db *sql.DB) {
	preparedStatement, err := db.Prepare(`
		INSERT INTO 
			daily_passive_income(date, source, asset_class, total, loss, net)
			VALUES(?,?,?,?,?,?)
	`)
	utils.HandleError(err)
	defer preparedStatement.Close()

	res, _ := preparedStatement.Exec(
		values.Date,
		values.Source,
		values.AssetClass,
		values.Total,
		values.Loss,
		values.Net)

	affectedRows, _ := res.RowsAffected()

	log.Println("inserted to daily_interests", values)
	log.Println("rows affected", affectedRows)
}

func InsertPortfolioValues(values PortfolioValueEntry, db *sql.DB) {
	preparedStatement, err := db.Prepare(`
		INSERT INTO 
			portfolio_values(date, source, value, initial_investment, profit, cash)
			VALUES(?,?,?,?,?,?)
	`)
	utils.HandleError(err)
	defer preparedStatement.Close()

	res, _ := preparedStatement.Exec(
		values.Date,
		values.Source,
		values.Value,
		values.InitialInvestment,
		values.Profit,
		values.Cash)

	affectedRows, _ := res.RowsAffected()

	log.Println("inserted to portfolio_values", values)
	log.Println("rows affected", affectedRows)
}

type Entry struct {
	Date       string
	Source     string
	AssetClass string
	Total      float64
	Loss       float64
	Net        float64
}

type PortfolioValueEntry struct {
	Date              string
	Source            string
	Value             float64
	InitialInvestment float64
	Cash              float64
	Profit            float64
}
