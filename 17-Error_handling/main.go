package main

import (
	"errors"
	"fmt"
)

func main() {
	defer fmt.Println("See you later")

	var x, y int

	fmt.Printf("Pleas enter first number:")
	fmt.Scan(&x)
	fmt.Printf("Pleas enter second number:")
	fmt.Scan(&y)

	result, err := divide(x, y)
	if err != nil {
		defer fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}

func divide(a, b int) (int, error) {
	var err error = nil
	var result int

	if b == 0 {
		err = errors.New("Any number cannot be divide by zero")
		panic("Zero division error")
	} else {
		result = a / b
	}

	return result, err
}
