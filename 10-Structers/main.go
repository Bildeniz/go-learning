package main

import "fmt"

func main() {
	var customer1 = Customer{id: 1, name: "Bilal", surname: "Deniz", age: 19}

	customer1.printall()
	customer1.changeName("Muhammed Bilal")
	customer1.printall()

	var customer2 Customer
	customer2.New(14, "Mahmut", "YILMAZ", 18)
	customer2.printall()
}

type Customer struct {
	id      int64
	name    string
	surname string
	age     int8
}

func (customer Customer) printall() {
	fmt.Printf("customer id: %d, name: %s, surname: %s, age: %d\n",
		customer.id, customer.name, customer.surname, customer.age)
}

func (customer *Customer) changeName(newName string) {
	customer.name = newName
}

func (customer *Customer) New(id int64, name string, surname string, age int8) {
	customer.id = id
	customer.name = name
	customer.surname = surname
	customer.age = age
}
