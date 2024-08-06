package main

import "fmt"

func main() {
	fmt.Println(add(15, 20))
	fmt.Println(calculation(20, 15))
	fmt.Println(summary(1, 2, 3, 4, 5))
}

func add(x int, y int) int {
	return x + y
}

func calculation(x int, y int) (int, int) {
	return x + y, x - y
}

func sum(numbers []int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func summary(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
