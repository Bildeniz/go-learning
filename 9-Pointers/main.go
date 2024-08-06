package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}

	fmt.Println(numbers)
	change(numbers)
	fmt.Println(numbers)
}

func change(numbers []int) {
	numbers[0] = 1000
}

func add12(a int) { // Passing by value
	a += 12
}

func add12pointer(a *int) { // Passing by variable
	*a = *a + 12
}
