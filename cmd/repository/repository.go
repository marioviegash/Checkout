package repository

import (
	"Checkout/cmd/model"
)

func Init() {
	model.Init()
	model.Item.Products["120P90"] = model.Product{
		Detail: model.ProductDetail{
			SKU: "120P90",
			Name: "Google Home",
			Price: 49.99,
		},
		Stock: 10,
	}

	model.Item.Products["43N23P"] = model.Product{
		Detail: model.ProductDetail{
			SKU: "43N23P",
			Name: "MacBook Pro",
			Price: 5399.99,
		},
		Stock: 5,
	}

	model.Item.Products["A304SD"] = model.Product{
		Detail: model.ProductDetail{
			SKU: "A304SD",
			Name: "Alexa Speaker",
			Price: 109.50,
		},
		Stock: 10,
	}

	model.Item.Products["234234"] = model.Product{
		Detail: model.ProductDetail{
			SKU: "234234",
			Name: "Rasberry Pi B",
			Price: 30.00,
		},
		Stock: 2,
	}
}