package repository

import (
	"context"
	"rinha/internal/database"
	"sync"
)

var cacheTable = sync.Map{}

func ClientExists(clientID int) (bool, error) {
	cached, found := cacheTable.Load(clientID)
	if found {
		return cached.(bool), nil
	}

	ctx := context.Background()

	db, err := database.Connection()
	if err != nil {
		return false, err
	}

	var exists int
	err = db.
		QueryRow(ctx, "SELECT id FROM clientes WHERE id = $1", clientID).
		Scan(&exists)
	if err != nil {
		return false, err
	}

	cacheTable.Store(clientID, exists > 0)
	return exists > 0, nil
}
