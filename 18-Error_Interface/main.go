package main

import "fmt"

type Product struct {
	name  string
	price float32
}

func (p *Product) New(name string, price float32) error {
	if len(name) == 0 {
		return NewError("Name cannot be null", 1)
	}
	if price <= 0 {
		return NewError("Price cannot be less than 0", 2)
	}

	p.name = name
	p.price = price

	return nil
}

func main() {
	var phone Product

	err := phone.New("Xiaomi Redmi Note 8 Pro", -42)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(phone)
	}

}

type ProductError struct {
	description string
	code        int
}

func (ProductError ProductError) Error() string {
	return fmt.Sprintf("Error description: %s, error code: %d", ProductError.description, ProductError.code)
}

func NewError(description string, code int) ProductError {
	return ProductError{description: description, code: code}
}
