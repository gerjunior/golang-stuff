package main

import "fmt"

type item struct {
	productID int
	quantity  int
	price     float64
}

type order struct {
	userID int
	items  []item
}

func (o order) totalValue() float64 {
	total := 0.0

	for _, item := range o.items {
		total += item.price * float64(item.quantity)
	}
	return total
}

func main() {
	order1 := order{
		userID: 1,
		items: []item{
			{
				productID: 123,
				quantity:  3,
				price:     12.10,
			},
			{
				productID: 111,
				quantity:  1,
				price:     100.00,
			},
			{
				productID: 12,
				quantity:  5,
				price:     20.99,
			},
		},
	}

	fmt.Println(order1.totalValue())
}
