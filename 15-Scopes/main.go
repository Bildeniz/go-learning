package main

import "fmt"

var x = 10

func main() {
	fmt.Println("x:", x)
	test()
	fmt.Println("x:", x)
}

func test() {
	x = 20
	fmt.Println("x:", x)
}
