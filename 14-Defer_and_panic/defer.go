package main

import "fmt"

func main() {
	defer fmt.Println("Defer running after your function")
	defer fmt.Println("Defer is running with stack")
	defer fmt.Println("Defer is a keyword")

	// No matters this part
	x := 100
	y := 200
	fmt.Printf("sum: %d\n", x+y)
}
