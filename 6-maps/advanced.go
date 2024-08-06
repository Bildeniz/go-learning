package main

import "fmt"

func main() {
	var peoples map[string][]string

	peoples = make(map[string][]string, 0)

	peoples["names"] = []string{"Bildeniz", "foo", "bar"}
	peoples["ages"] = []string{"19", "35", "32"}

	fmt.Println(peoples)

}
