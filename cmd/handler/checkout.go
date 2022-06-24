package handler

import (
	"Checkout/cmd/service"
	"encoding/json"
	"net/http"
)

type Response struct {
	Data Total
	Code int
	Status string
}

type Total struct {
	Total float32
}

func Checkout(w http.ResponseWriter, r * http.Request) {
	var c = service.New()
	c.AddItem("234234")
	c.AddItem("43N23P")
	var total = c.ScannedItemsTotalPrices()
	discount, _ := c.ApplyPromotions()
	total -= discount

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{
		Data: Total{ Total: total }, Code: 200, Status: "Success",
	})
}