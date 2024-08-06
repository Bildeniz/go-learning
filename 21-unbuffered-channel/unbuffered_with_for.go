package main

import (
	"fmt"
	"time"
)

func main() {
	myChannel := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			myChannel <- i
			fmt.Println("Writed value: ", i)
			time.Sleep(time.Second)
		}
		close(myChannel)
	}()

	func() {
		for {
			value, isOpen := <-myChannel
			if !isOpen {
				break
			}
			fmt.Println("Readed value: ", value)
		}
	}()

}
