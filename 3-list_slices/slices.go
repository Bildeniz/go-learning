package main

import (
	"fmt"
)

func main() {
	fruits := []string{
		"Apple", "Banana", "Orange", "Pear",
	}

	fmt.Println(fruits)

	fruits = append(fruits, "Watermelon")
	fmt.Println(fruits)

	fruits = append(fruits[:2], fruits[3:]...)
	fmt.Println(fruits)
}
