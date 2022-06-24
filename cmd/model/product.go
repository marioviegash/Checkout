package model

type ItemData struct {
	Products map[string]Product
}

type Product struct {
	Detail ProductDetail
	Stock int
}

type ProductDetail struct {
	SKU string
	Name string
	Price float32
}

var Item ItemData

func Init() {
	Item = ItemData {
		Products: make(map[string]Product),
	}
}