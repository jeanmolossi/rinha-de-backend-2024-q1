package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rinha/cmd/pkg"
	"rinha/internal/repository"
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	Value       int       `json:"valor"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	CreatedAt   time.Time `json:"realizada_em"`
}

func HandleNewTransaction(w http.ResponseWriter, r *http.Request) {
	// error can be safely ignored here
	// a previous validation was performed by middleware
	clientID, _ := strconv.Atoi(r.PathValue("id"))

	transaction := Transaction{}
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, pkg.JSONErr("can not read body", err), http.StatusBadRequest)
		return
	}

	if err := isBodyValid(transaction); err != nil {
		http.Error(w, pkg.JSONErr("invalid request received", err), http.StatusUnprocessableEntity)
		return
	}

	t, err := repository.NewTransaction(clientID, rune(transaction.Type[0]), transaction.Value, transaction.Description)
	if err != nil {
		http.Error(w, pkg.JSONErr("can not perform this transaction", err), http.StatusUnprocessableEntity)
		return
	}

	err = json.NewEncoder(w).Encode(t)
	if err != nil {
		http.Error(w, pkg.JSONErr("what was happen here ?", err), http.StatusInternalServerError)
	}
}

func isBodyValid(body Transaction) error {
	descLen := len(body.Description)
	if err := minLength(descLen, 1); err != nil {
		return err
	}

	if err := maxLength(descLen, 10); err != nil {
		return err
	}

	if err := enum(rune(body.Type[0]), []rune{'c', 'd'}); err != nil {
		return err
	}

	return nil
}

func minLength(l, min int) (err error) {
	if l < min {
		err = fmt.Errorf("field has %d but is required at least %d characters", l, min)
	}

	return
}

func maxLength(l, max int) (err error) {
	if l > max {
		err = fmt.Errorf("field has %d but is required at most %d characters", l, max)
	}

	return
}

func enum(val rune, allow []rune) error {
	for _, allowed := range allow {
		if val == allowed {
			return nil
		}
	}

	oneOf := make([]string, len(allow))
	for i, one := range allow {
		oneOf[i] = string(one)
	}

	return fmt.Errorf("field value is %s but should be one of %s", string(val), strings.Join(oneOf, ","))
}
