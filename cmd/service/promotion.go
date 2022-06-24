package service

import (
	"Checkout/cmd/model"
	"fmt"
	"strings"
)

func (c Cart) ApplyPromotions() (float32, string) {
	var discount float32
	var messages []string

	discount += c.XForPriceOfY(3, 2, "120P90")
	discount += c.OverXDiscountY(3, 10, "A304SD")

	d, m := c.BuyXGetYFree("43N23P", "234234")
	messages = append(messages, m)
	discount += d

	return discount, strings.Join(messages, "\n")
}

func (c Cart) XForPriceOfY(x, y int, id string) float32 {
	var discount float32

	if m, ok := c.Items[id]; ok {
		discount = model.Item.Products[id].Detail.Price * float32(m/x)
	}

	return discount
}

func (c Cart) OverXDiscountY(x int, y float32, id string) float32 {
	var discount float32

	if m, ok := c.Items[id]; ok {
		if m > x {
			discount = model.Item.Products[id].Detail.Price * float32(m) * (y / 100)
		}
	}

	return discount
}

func (c Cart) BuyXGetYFree(x, y string) (float32, string) {
	var count int
	var discount float32
	var additional int

	if xcount, ok := c.Items[x]; ok {
		additional = xcount
		if ycount, ok := c.Items[y]; ok {
			if ycount >= xcount {
				discount = model.Item.Products[y].Detail.Price * float32(xcount)
			}
			if ycount < xcount {
				discount = model.Item.Products[y].Detail.Price * float32(ycount)
				additional = xcount - ycount
			}
		}

		for i := 0; i < additional; i++ {
			err := c.AddItem(y)
			if err != nil {
				return discount, fmt.Sprintf("You have recieved %d free %s(s) along with your purchase of %s\nUnfortunatly we are out of stock to add the remaining free items", i, model.Item.Products[y].Detail.Name, model.Item.Products[x].Detail.Name)
			}
			count++
		}

	}

	return discount, fmt.Sprintf("You have recieved %d free %s(s) along with your purchase of %s", count, model.Item.Products[y].Detail.Name, model.Item.Products[x].Detail.Name)
}