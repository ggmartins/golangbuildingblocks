package minibank

import (
	"encoding/json"
	"fmt"
	auth "minibank/auth"
	datalayer "minibank/datalayer"
	models "minibank/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Web struct {
	dl   *datalayer.DataLayer_psql
	auth *auth.Auth
}

func (web *Web) SetDataLayer(dl *datalayer.DataLayer_psql) {
	web.dl = dl
}

func (web *Web) SetAuthorizer(auth *auth.Auth) {
	web.auth = auth //auth struct is currently empty
}

func (web *Web) Login(w http.ResponseWriter, r *http.Request) {
	var l models.Login
	err := json.NewDecoder(r.Body).Decode(&l)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"result":"Negado","error":%q}`, err)
		return
	} else {
		hashedSecret, err := web.dl.GetSecretHash(l)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, `{"result":"Negado","error":%q}`, err)
			return
		}
		if auth.CheckPasswordHash(l.Secret, hashedSecret) {
			userid, err := strconv.ParseUint(l.Cpf, 10, 64)
			if err == nil {
				token, err := auth.CreateToken(userid)
				if err == nil {
					w.WriteHeader(http.StatusOK)
					fmt.Fprintf(w, `{"result":"Aprovado","error":"","token":%q}`, token)
					return
				}
			}
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"result":"Negado","error":"Senha Invalida"}`)
	}
}

func (web *Web) GetAccountsIDBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := ExtractToken(r)
	isVerified, _ := auth.VerifyToken(token)
	if !isVerified {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"result":"Negado","error":"Desautorizado"}`)
		return
	}
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	balance, err := web.dl.GetAccountsIDBalance(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"result":"Negado","error":%q}`, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"result":"Aprovado","error":"","balance":%f}`, balance)
}

func (web *Web) PostAccounts(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"result":"Aprovado","error":""}`)
	}
}

func (web *Web) GetAccounts(w http.ResponseWriter, r *http.Request) {
	token := ExtractToken(r)
	isVerified, _ := auth.VerifyToken(token)
	if !isVerified {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"result":"Negado","error":"Desautorizado"}`)
		return
	}
	web.dl.PrintAccounts(w)
}

func (web *Web) GetTransfers(w http.ResponseWriter, r *http.Request) {
	token := ExtractToken(r)
	isVerified, _ := auth.VerifyToken(token)
	if !isVerified {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"result":"Negado","error":"Desautorizado"}`)
		return
	}
	web.dl.PrintTransfers(w)
}
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (web *Web) PostTransfers(w http.ResponseWriter, r *http.Request) {
	var t models.Transfer
	token := ExtractToken(r)
	isVerified, err := auth.VerifyToken(token)
	if !isVerified {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"result":"Negado","error":"Desautorizado"}`)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&t)
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
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"result":"Aprovado","error":""}`)
	}

}
