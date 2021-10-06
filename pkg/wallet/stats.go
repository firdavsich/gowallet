package wallet

import "database/sql"

// create new wallet
func Stats(db *sql.DB, id int) (transactions, summ int, err error) {
	// TODO select counted rows and sums
	err = db.QueryRow("SELECT COUNT(DISTINCT id), SUM(summ) FROM transactions WHERE wallet_id = $1 GROUP BY wallet_id", id).Scan(&transactions, &summ)
	return
}
