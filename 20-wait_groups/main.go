package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Func 1 is finished")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Func 2 is finished")
	}()

	wg.Wait()
	fmt.Println("Main function is finished")

}
