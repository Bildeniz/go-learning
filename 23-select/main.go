package main

import (
	"fmt"
	"time"
)

func main() {
	myChannel1 := make(chan string, 1)
	myChannel2 := make(chan string, 1)
	myChannel3 := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		myChannel1 <- "Bil"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		myChannel2 <- "den"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		myChannel3 <- "iz"
	}()

	for i := 1; i <= 3; i++ {
		select {
		case message := <-myChannel1:
			fmt.Printf("%s", message)
		case message := <-myChannel2:
			fmt.Printf("%s", message)
		case message := <-myChannel3:
			fmt.Printf("%s\n", message)
		}
	}

	fmt.Println("End of main")
}
