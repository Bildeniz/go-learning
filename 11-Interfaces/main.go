package main

import "fmt"

type IShippible interface {
	calculateShippingCost() int
}

type Books struct {
	size int
}

func (b Books) calculateShippingCost() int {
	return 5 + b.size*2
}

type Electronics struct {
	size int
}

func (e Electronics) calculateShippingCost() int {
	return 10 + e.size*4
}

func main() {
	shippibles := []IShippible{
		&Books{3},
		&Books{5},
		&Electronics{2},
		&Electronics{1},
	}

	total := 0

	for _, product := range shippibles {
		total += product.calculateShippingCost()
	}
	fmt.Println(total)

}
