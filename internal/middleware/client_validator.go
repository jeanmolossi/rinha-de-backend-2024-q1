package middleware

import (
	"net/http"
	"rinha/cmd/pkg"
	"rinha/internal/repository"
	"strconv"
)

func ClientValidator(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, pkg.JSONErr("client not found", nil), http.StatusNotFound)
			return
		}

		exists, err := repository.ClientExists(clientID)
		if err != nil {
			http.Error(w, pkg.JSONErr("client not found", err), http.StatusNotFound)
			return
		}

		if !exists {
			http.Error(w, pkg.JSONErr("client not found", nil), http.StatusNotFound)
			return
		}

		fn(w, r)
	}
}
