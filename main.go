package main

import (
	"Checkout/cmd/repository"
	"Checkout/cmd/router"
	"log"
	"net/http"
)

func main() {	
	log.Println("Seed all product data")
	repository.Init()

	log.Println("Set router")
	mux := http.NewServeMux()
	router.Init(mux)

	log.Println("Starting web checkout on port 8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
