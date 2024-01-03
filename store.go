package main

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "quotes.db"

func connectDB() {
	DB, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	db = DB

}

func createQuoteTableIfNotExists() {
	// create quotes table if it doesn't exist
	createQuoteTableSQL := `
		CREATE TABLE if NOT EXISTS quotes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			csvId INTEGER,
			quoteType INTEGER,
			content TEXT
		);
	`

	_, err := db.Exec(createQuoteTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

// parse a csv file with format: S/N | TipType | Content
func parseCSVFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, nil
}

// grab the last saved item from sqlite3
func getLastSavedCSVId() (int, error) {
	var lastCSVId int
	err := db.QueryRow("SELECT MAX(csvId) FROM quotes").Scan(&lastCSVId)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return lastCSVId, nil
}

// check if there are new items after last selected item
// collect all new items and store them in sqlite3 database
func insertQuotesFromCSV(filePath string) error {
	lines, err := parseCSVFile(filePath)
	if err != nil {
		return err
	}

	lastCSVId, err := getLastSavedCSVId()
	if err != nil {
		return err
	}

	for _, line := range lines {
		csvId, err := strconv.Atoi(line[0])
		if err != nil {
			return err
		}

		if csvId > lastCSVId {
			quoteType, err := strconv.Atoi(line[1])
			if err != nil {
				return err
			}

			content := line[2]

			_, err = db.Exec("INSERT INTO quotes (csvId, quoteType, content) VALUES (?, ?, ?)", csvId, quoteType, content)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
