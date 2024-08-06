package main

import "fmt"

func main() {
	var a, b int

	fmt.Printf("Pleas enter a number: ")
	_, err := fmt.Scanf("%d\n", &a)
	if err != nil {
		panic(err)
	}
	err = nil

	fmt.Printf("Pleas enter a number: ")
	_, err = fmt.Scanf("%d\n", &b)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result:", a+b)
}
