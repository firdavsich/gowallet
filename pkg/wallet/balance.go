package wallet

import "database/sql"

// create new wallet
func Balance(db *sql.DB, id int) (balance int, err error) {
	err = db.QueryRow("SELECT balance FROM wallets WHERE id = $1", id).Scan(&balance)
	if err != nil {
		return
	}
	return
}
