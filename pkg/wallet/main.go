package wallet

import (
	"database/sql"
	"fmt"
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

// TopUp wallet
func TopUp(db *sql.DB, id, summ int) (balance int, err error) {
	// check balance limit for identify <= 10000 and anonimous <=100000
	limitIdentified := 100000
	limitAnonimous := 10000

	var identified *bool
	err = db.QueryRow("SELECT verified,balance  FROM wallets WHERE id = $1", id).Scan(&identified, &balance)

	if err != nil {
		log.Println(err)
		return
	}

	finalBalance := balance + summ
	if finalBalance > limitAnonimous {
		if !*identified || finalBalance > limitIdentified {
			err = fmt.Errorf("limit is exceeded")
			log.Println(err)
			return
		}
	}

	// topUp wallet
	// sql code
	err = db.QueryRow("UPDATE wallets SET balance = $2 WHERE id = $1", id, finalBalance).Err()
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("ID=%d, Balance=%d, Identified=%v", id, balance, *identified)
	return finalBalance, err
}
