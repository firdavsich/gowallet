package wallet

import (
	"database/sql"
	"log"
)

// check wallet exist
func Check(db *sql.DB, id int) bool {
	var result int
	err := db.QueryRow("SELECT id FROM wallets WHERE id = $1", id).Scan(&result)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
