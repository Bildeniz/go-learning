package main

import (
	"fmt"
	"time"
)

func main() {
	values := make(chan int, 4) // Second parmeter is value size

	go func() {
		for i := 0; i <= 10; i++ {
			values <- i
			fmt.Println("Send data: ", i)
		}
		close(values)
	}()

	time.Sleep(2 * time.Second)
	for data := range values {
		fmt.Println("Readed value: ", data)
		time.Sleep(time.Second)
	}

	fmt.Println("End of the main")
}
