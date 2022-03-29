package minibank

import (
	"encoding/json"
	"fmt"
	auth "minibank/auth"
	datalayer "minibank/datalayer"
	models "minibank/models"
	"net/http"
)

type Web struct {
	dl   *datalayer.DataLayer_psql
	auth *auth.Auth
}

func (web *Web) SetDataLayer(dl *datalayer.DataLayer_psql) {
	web.dl = dl
}

func (web *Web) SetAuthorizer(auth *auth.Auth) {
	web.auth = auth
}

func (web *Web) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login\n")
}

func (web *Web) GetAccountsIDBalance(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetAccountsIDBalance\n")
}
func (web *Web) PostAccounts(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "PostAccounts\n")
	var a models.Account
	err := json.NewDecoder(r.Body).Decode(&a)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"result":"Negado","error":%q}`, err)
		return
	}
	a, err = auth.HashAccount(&a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"result":"Negado","error":%q}`, err)
		return
	}
	err = web.dl.InsertAccount(a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"result":"Negado","error":%q}`, err)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"result":"Negado","error":%q}`, err)
	} else {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, `{"result":"Aprovado","error":""}`)
	}
}

func (web *Web) GetAccounts(w http.ResponseWriter, r *http.Request) {
	web.dl.PrintAccounts(w)
}

func (web *Web) GetTransfers(w http.ResponseWriter, r *http.Request) {
	web.dl.PrintTransfers(w)
}

func (web *Web) PostTransfers(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "PostTransfers\n")
	var t models.Transfer
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err, status := web.dl.DoTransfer(&t)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"result":"%s","error":"Conta de origem não identificada (%q)"}`, err, status)
	} else {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, `{"result":"Aprovado","error":""}`)
	}

}
