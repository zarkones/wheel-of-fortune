package main

import (
	"api/ctrl"
	"net/http"
)

// initRoutes defines all the HTTP endpoints on our API.
func initRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /v1/spin/{userEmail}", ctrl.Spin)

	mux.HandleFunc("GET /v1/prizes/{spinID}", ctrl.GetPrize)

	mux.HandleFunc("PUT /v1/users", ctrl.Register)
}
