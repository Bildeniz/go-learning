package main

import (
	"context"
	"fmt"
	"time"
)

type Product struct {
	Id   int64
	Name string
}

var productChannel = make(chan Product)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	go getProduct(12)

	select {
	case p := <-productChannel:
		fmt.Println("Product id:", p.Id, "Name:", p.Name)
	case <-ctx.Done():
		fmt.Println(fmt.Errorf("Error, timeout occured"))
	}

	fmt.Println("End of the line")
}

func getProduct(id int64) {
	time.Sleep(time.Second * 4)
	productChannel <- Product{
		Id:   id,
		Name: "Laptop",
	}
}
