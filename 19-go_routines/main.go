package main

import (
	"fmt"
	"time"
)

func f1() {
	fmt.Println("function 1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("Function 1 finished")
}

func main() {
	fmt.Println("Main started")
	go f1()
	time.Sleep(5 * time.Second)

	fmt.Println("Main finished")

}
