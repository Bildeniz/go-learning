package main

import "fmt"

func main() {
	var names = map[string]string{
		"Bilal": "DENIZ",
		"foo":   "bar",
		"bla":   "bla",
	}

	for key, value := range names {
		fmt.Println(key, value)
	}
}
