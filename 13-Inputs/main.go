package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Welcome to calculator")
	fmt.Print("Pleas enter a number:")
	fmt.Scan(&a)
	fmt.Print("Pleas enter a number to:")
	fmt.Scan(&b)

	fmt.Printf("Result: %d\n", a+b)
}
