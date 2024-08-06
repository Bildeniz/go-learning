package main

import "fmt"

func main() {
	//var age int = 19
	//
	//if age >= 18 {
	//	fmt.Println("You can vote!")
	//} else {
	//	fmt.Println("You cant vote!")
	//}

	a := 15
	b := 16
	c := 17

	if a > b && a > c {
		fmt.Println("a is bigger than whole variables")
	} else if b > a && b > c {
		fmt.Println("b is bigger than whole variables")
	} else if c > a && c > b {
		fmt.Println("c is bigger than whole variables")
	}
}
