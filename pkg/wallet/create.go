package wallet

import "database/sql"

// create new wallet
func Create(db *sql.DB) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO wallets (verified,balance) VALUES (false,0) RETURNING id").Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, err
}
