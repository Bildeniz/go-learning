package main

import "fmt"

func main() {
	fruits := [4]string{
		"Apple", "Orange", "Banana", "Pear",
	}

	fmt.Println(fruits)
	//fruist = append(fruits, "Watermelon") There is a error because this is a fixed list
}
