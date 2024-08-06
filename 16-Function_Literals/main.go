package main

import "fmt"

func main() {
	var x, y int

	fmt.Printf("Lütfen Bir sayı giriniz:")
	fmt.Scan(&x)
	fmt.Printf("Lütfen ikinci sayıyı giriniz:")
	fmt.Scan(&y)

	sum := func(x, y int) int { return x + y }
	multiply := func(x, y int) int { return x * y }

	a, b := calculator(x, y, sum, multiply)

	fmt.Println(a, b)
}

func calculator(x int, y int, f1 func(int, int) int, f2 func(int, int) int) (int, int) {
	return f1(x, y), f2(x, y)
}
