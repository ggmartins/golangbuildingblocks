package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	//log "github.com/sirupsen/logrus"
)

/*const (
	host     = "postgres"
	port     = 5432
	user     = "minibank"
	password = "minibank"
	dbname   = "minibank"
)*/

type Account struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Cpf        string    `json:"cpf"`
	Secret     string    `json:"secret"`
	Balance    float64   `json:"balance"`
	Created_at time.Time `json:"created_at"`
}

type Transfer struct {
	Id         int       `json:"id"`
	IdSrc      int       `json:"account_origin_id"`
	IdDst      int       `json:"account_destination_id"`
	Amount     float64   `json:"amount"`
	Created_at time.Time `json:"created_at"`
}

var db *sql.DB

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login\n")
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := db.Query("SELECT * FROM accounts")
	if err != nil {
		panic(err)
	}
	for accounts.Next() {
		var row Account
		if err := accounts.Scan(&row.Id, &row.Name, &row.Cpf,
			&row.Secret, &row.Balance, &row.Created_at); err == nil {
			jsonOut, err := json.Marshal(row)
			if err != nil {
				log.Printf("error: %s\n", err)
			}
			fmt.Fprintf(w, "%s\n", string(jsonOut))
		} else {
			log.Println(err)
		}
	}
}

func GetAccountsIDBalance(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetAccountsIDBalance\n")
}
func PostAccounts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PostAccounts\n")
}
func GetTransfers(w http.ResponseWriter, r *http.Request) {
	accounts, err := db.Query("SELECT * FROM transfers")
	if err != nil {
		panic(err)
	}
	for accounts.Next() {
		var row Transfer
		if err := accounts.Scan(&row.Id, &row.IdSrc, &row.IdDst,
			&row.Amount, &row.Created_at); err == nil {
			jsonOut, err := json.Marshal(row)
			if err != nil {
				log.Printf("error: %s\n", err)
			}
			fmt.Fprintf(w, "%s\n", string(jsonOut))
		} else {
			log.Println(err)
		}
	}
}
func PostTransfers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PostTransfers\n")
}
func main() {
	var err error

	/*
		log.SetLevel(log.DebugLevel)

		formatter := &log.TextFormatter{
			FullTimestamp: true,
		}
		log.SetFormatter(formatter)
	*/

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"), os.Getenv("PGPORT"),
		os.Getenv("PGUSER"), os.Getenv("PGPASS"), os.Getenv("DBNAME"))
DBCONNECT:
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		time.Sleep(8 * time.Second)
		// Try to reconnect, server not up yet?
		// this is true on docker-compose env
		// maybe false for kube, so TODO:
		// let it die for kubernetes
		goto DBCONNECT
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println(err)
		time.Sleep(8 * time.Second)
		goto DBCONNECT
	}

	log.Println("Successfully db connected.")
	accounts, err := db.Query("SELECT * FROM accounts")
	if err != nil {
		panic(err)
	}
	for accounts.Next() {
		var row Account
		if err := accounts.Scan(&row.Id, &row.Name, &row.Cpf,
			&row.Secret, &row.Balance, &row.Created_at); err == nil {
			log.Printf("%d, %s, %s, %s, %f %s\n",
				row.Id, row.Name, row.Cpf,
				row.Secret, row.Balance, row.Created_at)
		} else {
			log.Println(err)
		}
	}

	r := mux.NewRouter()
	r.HandleFunc("/accounts", PostAccounts).Methods("POST")
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/accounts", GetAccounts).Methods("GET")
	r.HandleFunc("/accounts/{id}/balance", GetAccountsIDBalance).Methods("GET")
	r.HandleFunc("/transfers", GetTransfers).Methods("GET")
	r.HandleFunc("/transfers", PostTransfers).Methods("POST")
	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Web Server Started")
	http.ListenAndServe(":8000", r)
}
