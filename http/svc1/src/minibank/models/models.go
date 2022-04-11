package minibank

import "time"

type Account struct {
	Id         uint64    `json:"id"`
	Name       string    `json:"name"`
	Cpf        string    `json:"cpf"`
	Secret     string    `json:"secret"`
	Balance    float64   `json:"balance"`
	Created_at time.Time `json:"created_at"`
}

type Transfer struct {
	Id         uint64    `json:"id"`
	IdSrc      uint64    `json:"account_origin_id"`
	IdDst      uint64    `json:"account_destination_id"`
	Amount     float64   `json:"amount"`
	Created_at time.Time `json:"created_at"`
}

type Login struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}
