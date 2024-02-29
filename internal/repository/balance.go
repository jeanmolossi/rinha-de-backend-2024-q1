package repository

import (
	"context"
	"fmt"
	"rinha/internal/database"
	"time"

	"github.com/jackc/pgx/v5"
)

func GetUserBalance(clientID int) (Balance, error) {
	db, err := database.Connection()
	if err != nil {
		return Balance{}, err
	}

	b := new(pgx.Batch)

	b.Queue("SELECT c.valor, c.limite FROM saldos c WHERE c.cliente_id = $1", clientID)
	b.Queue("SELECT t.valor, t.tipo, t.descricao, t.realizada_em FROM transacoes t WHERE t.cliente_id = $1 ORDER BY t.realizada_em DESC LIMIT 10", clientID)

	ctx := context.Background()

	results := db.SendBatch(ctx, b)

	balance := Balance{
		Totals: Totals{
			BalanceTime: time.Now().Format(time.RFC3339),
		},
		LastTransactions: make([]Transaction, 0, 10),
	}

	row := results.QueryRow()
	err = row.Scan(&balance.Totals.Total, &balance.Totals.Limit)
	if err != nil {
		return Balance{}, err
	}

	rows, err := results.Query()
	if err != nil {
		return Balance{}, err
	}

	for rows.Next() {
		t := Transaction{}
		// t.valor, t.tipo, t.descricao, t.realizada_em
		err = rows.Scan(&t.Value, &t.Typ, &t.Description, &t.CreatedAt)
		if err != nil {
			fmt.Printf("append err %+v\n", err)
		}
		balance.LastTransactions = append(balance.LastTransactions, t)
	}

	results.Close()

	if rows.Err() != nil {
		return Balance{}, rows.Err()
	}

	return balance, nil
}
