package main

import "fmt"

func main() {
	for a := 0; a <= 10; a++ {
		if a%2 == 0 {
			continue
		}
		fmt.Println(a)
	}
}
