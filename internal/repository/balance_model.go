package repository

import "time"

type Balance struct {
	Totals           Totals        `json:"saldo"`
	LastTransactions []Transaction `json:"ultimas_transacoes"`
}

type Totals struct {
	Total       int    `json:"total"`
	BalanceTime string `json:"data_extrato"`
	Limit       int    `json:"limite"`
}

type Transaction struct {
	Value       int       `json:"valor"`
	Typ         string    `json:"tipo"`
	Description string    `json:"descricao"`
	CreatedAt   time.Time `json:"realizada_em"`
}
