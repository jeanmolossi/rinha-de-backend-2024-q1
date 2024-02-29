package api

import (
	"encoding/json"
	"net/http"
	"rinha/cmd/pkg"
	"rinha/internal/repository"
	"strconv"
)

func HandleRetrieveBalance(w http.ResponseWriter, r *http.Request) {
	clientID, _ := strconv.Atoi(r.PathValue("id"))

	balance, err := repository.GetUserBalance(clientID)
	if err != nil {
		http.Error(w, pkg.JSONErr("", err), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(balance)
	if err != nil {
		http.Error(w, "server failed", http.StatusInternalServerError)
		return
	}
}
