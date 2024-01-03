package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func main() {
	fmt.Println("Bot is up! 🚀")
	setUp()
	defer db.Close()
	// createTweet()
}
