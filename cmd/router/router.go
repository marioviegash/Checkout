package router

import (
	"Checkout/cmd/handler"
	"encoding/json"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Status string `json:"status"`
}

func Init(mux *http.ServeMux) {
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/v1/api", apiHandler)
	mux.HandleFunc("/v1/api/checkout", handler.Checkout)
}

func homeHandler(w http.ResponseWriter, r * http.Request) {
	w.Write([]byte("Welcome to Checkout Apps"))
}

func apiHandler(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{
		Code: 200, Status: "Success",
	})
}