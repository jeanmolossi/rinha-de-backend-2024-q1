package main

import (
	"net/http"
	"net/http/pprof"
	"rinha/internal/api"
	"rinha/internal/middleware"
)

func main() {
	mux := http.DefaultServeMux

	mdw := middleware.Compose(
		middleware.ClientValidator,
		middleware.LogRequest,
	)

	mux.HandleFunc("GET /debug/pprof/cpu", pprof.Profile)
	mux.HandleFunc("POST /clientes/{id}/transacoes", mdw(api.HandleNewTransaction))
	mux.HandleFunc("GET /clientes/{id}/extrato", mdw(api.HandleRetrieveBalance))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
