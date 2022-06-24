package service

import (
	"Checkout/cmd/model"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// NoItemsInCart is the fixed message for logging when a cart has zero items
const NoItemsInCart = "No items in shopping cart"

// OutOfStock is the fixed message for logging when an item is out of stock
const OutOfStock = "Out of stock"

// ZeroItemCountInCart is the fixed message for logging when a cart has an item with zero items added
const ZeroItemCountInCart = "Item with zero count in cart"

type Cart struct {
	Items map[string] int
}

func New() *Cart {
	return &Cart{
		Items: make(map[string]int),
	}
}

func (c Cart) AddItem(id string) error {
	_, ok := model.Item.Products[id]; 
	if !ok {
		return fmt.Errorf("Product does not exist in inventory: %s", id)
	}

	p := model.Item.Products[id]
	if p.Stock < 1 {
		return fmt.Errorf("Not enough stock of %s (%s)", p.Detail.Name, p.Detail.SKU)
	}

	p.Stock--
	model.Item.Products[id] = p

	var count = 1
	if _, ok := c.Items[id]; ok {
		count = c.Items[id] + 1
	}

	c.Items[id] = count

	return nil
}

func (c Cart) ScannedItemsLabels() string {
	if len(c.Items) == 0 {
		return NoItemsInCart
	}

	var items []string
	for id, count := range c.Items {
		if count == 0 {
			continue
		}

		if p, ok := model.Item.Products[id]; ok {
			if count > 1 {
				items = append(items, p.Detail.Name+" x"+strconv.Itoa(count))
			} else {
				items = append(items, p.Detail.Name)
			}

		} else {
			fmt.Errorf("Item missing from inventory")
		}
	}

	if len(items) == 0 {
		return NoItemsInCart
	}

	sort.Strings(items)
	return strings.Join(items, ", ")
}

func (c Cart) ScannedItemsTotalPrices() float32 {
	var total float32

	for id, count := range c.Items {
		if d, ok := model.Item.Products[id]; ok {
			total = total + (d.Detail.Price * float32(count))
		} else {
			fmt.Errorf("Item missing from inventory")
		}
	}

	return total
}