package repository

import (
	"context"
	"rinha/internal/database"
)

func NewTransaction(clientID int, txType rune, val int, d string) (AccountLimits, error) {
	var empty AccountLimits

	ctx := context.Background()

	db, err := database.Connection()
	if err != nil {
		return empty, err
	}

	// select * from atualiza_saldo(1, 100000, 'd', 'descricao');
	rows, err := db.Query(ctx,
		"SELECT * FROM atualiza_saldo($1, $2, $3, $4)",
		clientID, val, string(txType), d,
	)
	if err != nil {
		return empty, err
	}

	var limit, balance int

	for rows.Next() {
		if err := rows.Scan(&balance, &limit); err != nil {
			return empty, err
		}
	}

	if rows.Err() != nil {
		return empty, rows.Err()
	}

	return AccountLimits{
		Limit:   limit,
		Balance: balance,
	}, nil
}
