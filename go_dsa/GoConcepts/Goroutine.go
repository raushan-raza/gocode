package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	go printNumbers()                   // Start the goroutine
	time.Sleep(2000 * time.Millisecond) // Give the goroutine time to complete
	fmt.Println("Main function ends")
}
