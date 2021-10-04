package api

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/firdavsich/gowallet/pkg/wallet"
	_ "github.com/lib/pq"
)

var (
	port = os.Getenv("PORT")

	dbConn *sql.DB
)

func checkFunc(rw http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.RequestURI)
	id, _ := strconv.Atoi(r.FormValue("id"))
	if wallet.Check(dbConn, id) == true {
		fmt.Fprintf(rw, "exist")
	} else {
		fmt.Fprintf(rw, "not exist")

	}
}
func createFunc(rw http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.RequestURI)
	if id, err := wallet.Create(dbConn); err != nil {
		fmt.Fprintf(rw, "Error")

	} else {
		fmt.Fprintf(rw, strconv.Itoa(id))
	}
}

func Run() {
	var err error
	dbConn, err = sql.Open("postgres", dbConninfo)
	if err != nil {
		log.Panic(err)
	}

	http.HandleFunc("/check", checkFunc)
	http.HandleFunc("/create", createFunc)

	log.Printf("API server on %s", port)
	err = http.ListenAndServe(net.JoinHostPort("", port), nil)
	if err != nil {
		log.Fatal(err)
	}

}
