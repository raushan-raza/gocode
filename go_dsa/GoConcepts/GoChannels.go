package main

import (
	"fmt"
)

func sum(a, b int, c chan int) {
	c <- a + b // Send the sum to the channel
}

func main() {
	c := make(chan int) // Create a channel of integers
	go sum(1, 2, c)     // Start a goroutine and send result to channel
	go sum(3, 4, c)     // Start another goroutine and send result to channel
	x := <-c            // Receive from channel
	y := <-c            // Receive from channel
	fmt.Println(x, y)   // Print received values
}
