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

// create new wallet
func Create(db *sql.DB) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO wallets (verified,balance) VALUES (false,0) RETURNING id").Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, err
}
