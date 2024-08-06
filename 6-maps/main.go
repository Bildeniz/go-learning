package main

import "fmt"

func main() {
	var names map[string]int
	names = make(map[string]int, 0)

	names["Bilal"] = 19
	names["Foo"] = 42
	names["Bar"] = 45

	delete(names, "Bar")

	fmt.Println(names)
}
