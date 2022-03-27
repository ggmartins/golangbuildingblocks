package main

import (
	"log"
	auth "minibank/auth"
	datalayer "minibank/datalayer"
	handlers "minibank/webhandlers"

	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	//log "github.com/sirupsen/logrus"
)

func main() {

	webHandlers := handlers.Web{}
	dl := datalayer.DataLayer_psql{}
	auth := auth.Auth{}

	dl.Init()
	dl.PrintAccounts(nil)

	webHandlers.SetDataLayer(&dl)
	webHandlers.SetAuthorizer(&auth)

	r := mux.NewRouter()
	r.HandleFunc("/accounts", webHandlers.PostAccounts).Methods("POST")
	r.HandleFunc("/login", webHandlers.Login).Methods("POST")
	r.HandleFunc("/accounts", webHandlers.GetAccounts).Methods("GET")
	r.HandleFunc("/accounts/{id}/balance", webHandlers.GetAccountsIDBalance).Methods("GET")
	r.HandleFunc("/transfers", webHandlers.GetTransfers).Methods("GET")
	r.HandleFunc("/transfers", webHandlers.PostTransfers).Methods("POST")

	//don't need this for now
	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Web Server Started")
	http.ListenAndServe(":8000", r)
}
