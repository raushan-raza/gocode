package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // Simulate work
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2 // Send result to the results channel
	}
}

func main() {
	jobs := make(chan int, 5)    // Buffered channel for jobs
	results := make(chan int, 5) // Buffered channel for results

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results) // Start 3 worker goroutines
	}

	for j := 1; j <= 5; j++ {
		jobs <- j // Send 5 jobs to the jobs channel
	}
	close(jobs) // Close the jobs channel

	for a := 1; a <= 5; a++ {
		fmt.Println(<-results) // Receive all results from the results channel
	}
}
