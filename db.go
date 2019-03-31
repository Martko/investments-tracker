package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"investments-tracker/utils"
	"log"
)

func getDbConnection() (db *sql.DB) {
	dbUser, dbPass, dbName := getDatabaseCredentials()

	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@/"+dbName)
	utils.HandleError(err)

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	utils.HandleError(err)

	return db
}

func insertOrUpdateDatabase(values dbEntry) {
	db := getDbConnection()
	selectQuery := `
		SELECT
			id
	    FROM
	    	interests
	    WHERE
	    	source = ? AND
	    	month = ? AND 
	    	year = ? 
	    LIMIT 1`

	var id int
	err := db.QueryRow(selectQuery, values.source, values.month, values.year).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			insertValues(values, db)
		} else {
			utils.HandleError(err)
		}
	} else {
		updateValues(values, id, db)
	}

	defer db.Close()
}

func insertValues(values dbEntry, db *sql.DB) {
	preparedStatement, err := db.Prepare(`
		INSERT INTO 
			interests(source, month, year, interest_amount, loss_amount, net_profit)
			VALUES(?,?,?,?,?,?)
	`)
	utils.HandleError(err)

	preparedStatement.Exec(
		values.source,
		values.month,
		values.year,
		values.interestAmount,
		values.lossAmount,
		values.netProfit)

	log.Println("inserted to database", values)
}

func updateValues(values dbEntry, rowId int, db *sql.DB) {
	preparedStatement, err := db.Prepare(`
		UPDATE
			interests
		SET
			interest_amount=?,
			loss_amount=?,
			net_profit=?
		WHERE
			id=?
	`)

	utils.HandleError(err)

	preparedStatement.Exec(
		values.interestAmount,
		values.lossAmount,
		values.netProfit,
		rowId)

	utils.HandleError(err)
	log.Println("values updated", values)
}

type dbEntry struct {
	source         string
	month          int
	year           int
	interestAmount float64
	lossAmount     float64
	netProfit      float64
}
