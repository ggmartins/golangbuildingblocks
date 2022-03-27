package minibank

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	models "minibank/models"
	"os"
	"time"
)

type DataLayer_psql struct {
	db *sql.DB
	DataLayer
}

func (dl *DataLayer_psql) Init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"), os.Getenv("PGPORT"),
		os.Getenv("PGUSER"), os.Getenv("PGPASS"), os.Getenv("DBNAME"))
DBCONNECT:
	dl.db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		time.Sleep(8 * time.Second)
		// Try to reconnect, server not up yet?
		// this is true on docker-compose env
		// maybe false for kube, so TODO:
		// let it die for kubernetes
		goto DBCONNECT
	}

	err = dl.db.Ping()
	if err != nil {
		log.Println(err)
		time.Sleep(8 * time.Second)
		goto DBCONNECT
	}
	log.Println("Successfully db connected.")
}

func (dl *DataLayer_psql) Close() {
	dl.db.Close()
}

func (dl *DataLayer_psql) PrintAccounts(output io.Writer) error {
	accounts, err := dl.db.Query("SELECT * FROM accounts")
	if err != nil {
		panic(err)
	}
	for accounts.Next() {
		var row models.Account
		if err := accounts.Scan(&row.Id, &row.Name, &row.Cpf,
			&row.Secret, &row.Balance, &row.Created_at); err == nil {
			if output == nil {
				log.Printf("%d, %s, %s, %s, %f %s\n",
					row.Id, row.Name, row.Cpf,
					row.Secret, row.Balance, row.Created_at)
			} else {
				jsonOut, err := json.Marshal(row)
				if err != nil {
					log.Printf("error: %s\n", err)
					return err
				}
				fmt.Fprintf(output, "%s\n", string(jsonOut))
			}
		} else {
			log.Println(err)
			return err
		}
	}
	return err
}

func (dl *DataLayer_psql) PrintTransfers(output io.Writer) error {
	accounts, err := dl.db.Query("SELECT * FROM transfers")
	if err != nil {
		panic(err)
	}
	for accounts.Next() {
		var row models.Transfer
		if err := accounts.Scan(&row.Id, &row.IdSrc, &row.IdDst,
			&row.Amount, &row.Created_at); err == nil {
			if output == nil {
				//TODO log transfers
			} else {
				jsonOut, err := json.Marshal(row)
				if err != nil {
					log.Printf("error: %s\n", err)
					return err
				}
				fmt.Fprintf(output, "%s\n", string(jsonOut))
			}
		} else {
			log.Println(err)
		}
	}
	return err
}

func (dl *DataLayer_psql) DoTransfer(t *models.Transfer) (error, string) {
	var err error
	var IdSrc int
	var BalanceSrc float64
	err = dl.db.QueryRow("SELECT id, balance from accounts where id = $1", t.IdSrc).Scan(&IdSrc, &BalanceSrc)
	if err != nil {
		log.Printf(`{"result":"Negado","error":"Conta de origem não identificada (%q)"}`, err)
		return errors.New("Conta de origem não identificada"), "Negado"
	}

	var IdDst int
	var BalanceDst float64
	err = dl.db.QueryRow("SELECT id, balance from accounts where id = $1", t.IdDst).Scan(&IdDst, &BalanceDst)
	if err != nil {
		log.Printf(`{"result":"Negado","error":"Conta de destino não identificada %q"}`, err)
		return errors.New("Conta de destino não identificada"), "Negado"
	}
	log.Printf("*** IdSrc: %d\n", IdSrc)
	log.Printf("*** IdDst: %d\n", IdDst)
	log.Printf("*** Amount: %f\n", t.Amount)
	log.Printf("*** BalanceSrc: %f\n", BalanceSrc)
	log.Printf("*** BalanceDst: %f\n", BalanceDst)
	NewBalanceSrc := BalanceSrc - t.Amount
	NewBalanceDst := BalanceDst + t.Amount
	if NewBalanceSrc < 0.0 {
		log.Printf(`{"result":"Negado","error":"Saldo Insuficiente"}`)
		return errors.New("Saldo Insuficiente"), "Negado"
	} else {
		//TODO insert transfer here, use status field to set initial pending status
		log.Printf(`{"result":"Aprovado","error":""}`)
	}
	// lift https://www.sohamkamani.com/golang/sql-transactions/
	// Transactions need context, maybe because the need for op timeout cancellation (?)
	ctx := context.Background()
	tx, err := dl.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return errors.New("Internal Error (Context)"), "Negado"
	}
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = $1 WHERE id = $2", NewBalanceSrc, IdSrc)
	if err != nil {
		tx.Rollback()
		return errors.New("Internal Error (UPDATE 1)"), "Negado"
	}
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = $1 WHERE id = $2", NewBalanceDst, IdDst)
	if err != nil {
		tx.Rollback()
		return errors.New("Internal Error (UPDATE 2)"), "Negado"
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO transfers "+
		"(account_origin_id, account_destination_id, amount,  created_at) "+
		"VALUES ($1, $2, $3, $4)", IdSrc, IdDst, t.Amount, time.Now())
	if err != nil {
		tx.Rollback()
		return errors.New("Internal Error (INSERT)"), "Negado"
	}
	tx.Commit()

	ctx.Done()
	return nil, "Aprovado"
}
